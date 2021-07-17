package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"golang.org/x/net/http2"

	"backend/config"
	"backend/internal/pkg/logger"
	"backend/internal/pkg/middlewares"
)

type HTTP struct {
	server *http.Server
	Router *mux.Router
	mws    *middlewares.Middleware
}

func New(mws *middlewares.Middleware) *HTTP {
	addr := fmt.Sprintf(":%s", config.Cfg.HttpPort)

	router := mux.NewRouter()

	server := http.Server{
		Addr:    addr,
		Handler: mws.CorsMWHTTP(router),
	}

	return &HTTP{
		server: &server,
		Router: router,
		mws:    mws,
	}
}

func (h *HTTP) Run() {
	logger.Log.Println("HTTP REST server started on", h.server.Addr)
	//h.printRoutes()

	go func() {
		log.Fatal(h.server.ListenAndServe())
	}()
}

func (h *HTTP) RunTLS(certFile, keyFile string) {
	h.server.Addr = ":443"
	logger.Log.Println("HTTP REST server in TLS Mode started on", h.server.Addr)

	//h.printRoutes()

	// http to https redirection
	logger.Log.Println("HTTP to HTTPS Redirection server started on", "0.0.0.0:80")
	go func() {
		logger.Log.Fatal(http.ListenAndServe(":80", http.HandlerFunc(h.mws.HTTPSRedirectMW)))
	}()

	logger.Log.Fatal(h.server.ListenAndServeTLS(certFile, keyFile))
}

func (h *HTTP) setupHTTP2() {
	logger.Log.Fatal(http2.ConfigureServer(h.server, nil))
}

func (h *HTTP) printRoutes() {
	routesList := reflect.ValueOf(h.Router).Elem()

	fmt.Printf("routes: %v\n", routesList.FieldByName("m"))
}

func (h *HTTP) Shutdown(ctx context.Context) {
	logger.Log.Println("emergency termination call. terminating HTTP REST Web server")

	err := h.server.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	logger.Log.Println("HTTP REST server terminated")
}
