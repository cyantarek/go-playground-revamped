package grpcweb

import (
	"backend/config"
	"backend/internal/pkg/logger"
	"context"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type GRPCWeb struct {
	proxy *http.Server
}

func NewGRPCWeb(srv *grpc.Server) *GRPCWeb {
	gRPCWebServer := grpcweb.WrapServer(srv)

	addr := fmt.Sprintf("%s:%s", config.Cfg.GRPCWebHost, config.Cfg.GRPCWebPort)

	httpServer := &http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				log.Println("HTTP/2 Request arrived", r.Proto)
				gRPCWebServer.ServeHTTP(w, r)
			} else {
				log.Println("HTTP/1.1 Request arrived", r.Proto)
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")

				if gRPCWebServer.IsGrpcWebRequest(r) {
					log.Println("gRPC-Web Request arrived", r.Proto)
					gRPCWebServer.ServeHTTP(w, r)
				}
			}
		}),
	}

	return &GRPCWeb{proxy: httpServer}
}

func (g *GRPCWeb) Run() {
	log.Println("gRPC web server started on", g.proxy.Addr)
	go func() {
		log.Fatal(g.proxy.ListenAndServe())
	}()
}

func (g *GRPCWeb) Shutdown(ctx context.Context) {
	logger.Log.Println("emergency termination call. terminating gRPC server")

	err := g.proxy.Shutdown(ctx)
	if err != nil {
		log.Println(err)
	}

	logger.Log.Println("gRPC server terminated")
}
