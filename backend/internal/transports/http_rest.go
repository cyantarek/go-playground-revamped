package transports

import (
	"backend/api/playground"
	"backend/config"
	"backend/internal/middlewares"
	"fmt"
	"github.com/gogo/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type HttpTransport struct {
	mux  *http.ServeMux
	gMux *runtime.ServeMux
	srv  *http.Server
}

func (h *HttpTransport) Register(gRPCAddr string, opts []grpc.DialOption) error {
	return playground.RegisterPlaygroundHandlerFromEndpoint(context.Background(), h.gMux, gRPCAddr, opts)
}

func BaseHttpTransport(cfg *config.Config) *HttpTransport {
	mux := http.NewServeMux()

	jsonPb := &gateway.JSONPb{
		EmitDefaults: true,
		Indent:       "  ",
		OrigName:     true,
	}
	gMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, jsonPb))
	mux.Handle("/", gMux)

	restAddr := fmt.Sprintf("%s:%s", cfg.Server.REST.Host, cfg.Server.REST.Port)

	srv := http.Server{
		Addr: restAddr,
		//Handler: middlewares.Cors(middlewares.Auth(middlewares.Log(mux))), // req -> cors -> auth -> server -> log -> user
		Handler: middlewares.Cors(mux), // req -> cors -> auth -> server -> log -> user
	}

	return &HttpTransport{
		mux:  mux,
		gMux: gMux,
		srv:  &srv,
	}
}

func (h *HttpTransport) Run() {
	log.Println("HTTP REST server started")
	go func() {
		log.Fatal(h.srv.ListenAndServe())
	}()
}

func (h *HttpTransport) RunTLS(certFile, keyFile string) {
	log.Println("HTTP REST server in TLS Mode started")
	go func() {
		log.Fatal(h.srv.ListenAndServeTLS(certFile, keyFile))
	}()
}

func (h *HttpTransport) setupHttp2() {
	log.Fatal(http2.ConfigureServer(h.srv, nil))
}

func (h *HttpTransport) setupSwagger() {
	// routes for swagger
	fs := http.FileServer(http.Dir("third_party/swagger-ui/static"))
	h.mux.Handle("/static/", http.StripPrefix("/static/", fs))

	h.mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "third_party/swagger-ui/index.html")
	})

	h.mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/swagger/kickme.swagger.json.example")
	})
}
