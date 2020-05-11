package middlewares

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Accept,Authorization")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS,HEAD")
			return
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		h.ServeHTTP(w, r)
	})
}

// Get Response code from Response Writer recorder
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

func Log(h http.Handler) http.Handler {
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

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" {
			//s.logger.Warn("token not present")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func JwtInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	m, _ := metadata.FromIncomingContext(ctx)

	token := m.Get("authorization")

	if len(token) == 0 {
		return nil, status.Error(codes.Unauthenticated, "Unauthorized: token not found")
	}

	// if token present, validate it

	h, err := handler(ctx, req)
	if err != nil {
		//s.logger.WithField("err", err.Error()).Error("RPC failed with error")
	}
	return h, err
}

func HttpsRedirect(w http.ResponseWriter, r *http.Request) {
	// remove/add not default ports from req.Host
	target := "https://" + r.Host + r.URL.Path
	if len(r.URL.RawQuery) > 0 {
		target += "?" + r.URL.RawQuery
	}
	log.Printf("redirect to: %s", target)
	http.Redirect(w, r, target,
		// see comments below and consider the codes 308, 302, or 301
		http.StatusTemporaryRedirect)
}
