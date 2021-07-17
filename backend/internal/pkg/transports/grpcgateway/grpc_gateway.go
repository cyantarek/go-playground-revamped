package grpcgateway

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"reflect"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"

	"backend/config"
	"backend/internal/pkg/logger"
	"backend/internal/pkg/middlewares"
)

type GRPCGateway struct {
	gmux *runtime.ServeMux
	srv  *http.Server
}

func (g *GRPCGateway) Server() *runtime.ServeMux {
	return g.gmux
}

func NewGRPCGateway(mws *middlewares.Middleware) *GRPCGateway {
	gmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{}))

	srv := http.Server{
		Addr: fmt.Sprintf(":%s", config.Cfg.APIPort),
		Handler: mws.CorsMWHTTP(
			mws.AuthMWHTTP(
				mws.PanicRecoveryMWHTTP(
					mws.LogMWHTTP(gmux),
				),
			),
		),
	}

	return &GRPCGateway{
		gmux: gmux,
		srv:  &srv,
	}
}

func setupPprof(m *http.ServeMux) {
	m.HandleFunc("/debug/pprof/", func(w http.ResponseWriter, r *http.Request) {
		part := r.RequestURI[len("/debug/pprof/"):]

		if part == "" {
			pprof.Handler("index").ServeHTTP(w, r)
		} else {
			pprof.Handler(part).ServeHTTP(w, r)
		}
	})
}

func (g *GRPCGateway) Run() {
	logger.Log.Println("GRPC Gateway server started on", g.srv.Addr)
	g.printRoutes()

	go func() {
		log.Fatal(g.srv.ListenAndServe())
	}()
}

func (g *GRPCGateway) RunTLS(certFile, keyFile string) {
	g.srv.Addr = ":443"
	logger.Log.Println("GRPC Gateway server in TLS Mode started on", g.srv.Addr)
	//g.printRoutes()

	// http to https redirection
	/*logger.LogMWHTTP.Println("HTTP to HTTPS Redirection server started on", "0.0.0.0:80")
	go func() {
		logger.LogMWHTTP.Fatal(http.ListenAndServe(":80", http.HandlerFunc(middlewares.HttpsRedirect)))
	}()*/

	logger.Log.Fatal(g.srv.ListenAndServeTLS(certFile, keyFile))
}

func (g *GRPCGateway) printRoutes() {
	routesList := reflect.ValueOf(g.gmux).Elem()

	fmt.Println(routesList)
}

func (g *GRPCGateway) setupHttp2() {
	logger.Log.Fatal(http2.ConfigureServer(g.srv, nil))
}

func (g *GRPCGateway) Shutdown(ctx context.Context) {
	logger.Log.Println("emergency termination call. terminating HTTP REST Web server")

	err := g.srv.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	logger.Log.Println("HTTP REST server terminated")
}
