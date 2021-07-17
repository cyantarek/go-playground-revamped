package middlewares

import (
	"backend/config"
	"backend/pkg/jwtpacker"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

func (mw Middleware) CorsMWHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,Token,Authorization,X-User-Agent")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,HEAD")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		h.ServeHTTP(w, r)
	})
}

// GetOrder Response code from Response Writer recorder
// https://upgear.io/blog/golang-tip-wrapping-http-response-writer-for-middleware/
// https://ndersson.me/post/capturing_status_code_in_net_http/
type loggingStatusRecorder struct {
	http.ResponseWriter
	status int
}

func (l *loggingStatusRecorder) WriteHeader(code int) {
	l.status = code
	l.ResponseWriter.WriteHeader(code)
}

func (mw Middleware) LogMWHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loggingHandler := loggingStatusRecorder{w, 200}

		start := time.Now()

		var reqBodyCopy bytes.Buffer

		tee := io.TeeReader(r.Body, &reqBodyCopy)
		r.Body = ioutil.NopCloser(tee)

		h.ServeHTTP(&loggingHandler, r)

		end := time.Now()
		duration := end.Sub(start).Nanoseconds() / 1000000

		reqBody, _ := ioutil.ReadAll(&reqBodyCopy)
		reqBody = nil

		fields := logrus.Fields{
			"request": logrus.Fields{
				"host":        r.Host,
				"path":        r.URL.Path,
				"query":       r.URL.RawQuery,
				"method":      r.Method,
				"remote_addr": r.RemoteAddr,
				//"headers":     r.Header,
				"uri":         r.RequestURI,
				"form_values": r.PostForm,
				"body":        string(reqBody),
				"referer":     r.Referer(),
				"user_agent":  r.UserAgent(),
			},
			"response": logrus.Fields{
				"status": loggingHandler.status,
			},
			"timing": logrus.Fields{
				"start":        start.Format(time.RFC3339),
				"end":          end.Format(time.RFC3339),
				"duration(ms)": duration,
			},
		}

		fmt.Println(fields)

		//s.logger.WithFields(fields).Info("request handling completed")
	})
}

func (mw Middleware) PanicRecoveryMWHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err) // May be log this error? Send to sentry?

				jsonBody, _ := json.Marshal(map[string]string{
					"error": "There was an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}

		}()

		h.ServeHTTP(w, r)

	})
}

func (mw Middleware) AuthMWHTTP(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/static/") {
			h.ServeHTTP(w, r)
			return
		}

		if config.Cfg.AuthSkipper[r.URL.Path] {
			h.ServeHTTP(w, r)
			return
		}

		// extract token from header
		token := r.Header.Get("Authorization")

		if token == "" {
			cook, err := r.Cookie("token")
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token = cook.Value
		}

		tokenParts := strings.Split(token, " ")

		if len(tokenParts) != 2 {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if valid, claims := jwtpacker.ValidateToken(tokenParts[1]); valid {
			if email, ok := claims["email"]; ok {
				//r.Header.Set("Grpc-Metadata-Email", email.(string))
				//r.Header.Set("Grpc-Metadata-Token", tokenParts[1])

				ctx := context.WithValue(r.Context(), "email", email.(string))

				h.ServeHTTP(w, r.WithContext(ctx))
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	})
}

func (mw Middleware) HTTPSRedirectMW(w http.ResponseWriter, r *http.Request) {
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s\n", target)
	http.Redirect(w, r, target,
		http.StatusTemporaryRedirect)
}

func (mw Middleware) HttpRedirect(w http.ResponseWriter, r *http.Request) {
	target := "http://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s\n", target)
	http.Redirect(w, r, target,
		http.StatusTemporaryRedirect)
}

func (mw Middleware) RateLimitMWHTTP(h http.Handler) http.Handler {
	rateLimiter := newIPRateLimiter(1, 1)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		currentLimit := rateLimiter.GetLimiter(r.RemoteAddr)
		if !currentLimit.Allow() {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		h.ServeHTTP(w, r)
	})
}

type ipRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  *sync.RWMutex
	r   rate.Limit
	b   int
}

// newIPRateLimiter .
func newIPRateLimiter(r rate.Limit, b int) *ipRateLimiter {
	i := &ipRateLimiter{
		ips: make(map[string]*rate.Limiter),
		mu:  &sync.RWMutex{},
		r:   r,
		b:   b,
	}

	return i
}

// AddIP creates a new rate limiter and adds it to the ips map,
// using the IP address as the key
func (i *ipRateLimiter) AddIP(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()

	limiter := rate.NewLimiter(i.r, i.b)

	i.ips[ip] = limiter

	return limiter
}

// GetLimiter returns the rate limiter for the provided IP address if it exists.
// Otherwise calls AddIP to add IP address to the map
func (i *ipRateLimiter) GetLimiter(ip string) *rate.Limiter {
	i.mu.Lock()
	limiter, exists := i.ips[ip]

	if !exists {
		i.mu.Unlock()
		return i.AddIP(ip)
	}

	i.mu.Unlock()

	return limiter
}
