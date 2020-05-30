package transports

import (
	"backend/config"
	"fmt"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type GRPCWebTransport struct {
	proxy *http.Server
}

func BaseGRPCWebTransport(srv *grpc.Server, cfg *config.Config) *GRPCWebTransport {
	gRPCWebServer := grpcweb.WrapServer(srv)

	addr := fmt.Sprintf("%s:%s", cfg.Server.GrpcWeb.Host, cfg.Server.GrpcWeb.Port)

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

	return &GRPCWebTransport{proxy: httpServer}
}

func (transport *GRPCWebTransport) Run() {
	log.Println("gRPC Web server started on", transport.proxy.Addr)
	go func() {
		log.Fatal(transport.proxy.ListenAndServe())
	}()
}

func (transport *GRPCWebTransport) Shutdown() {
	log.Println("emergency termination call. terminating gRPC Web server")
	
	err := transport.proxy.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}
	
	log.Println("gRPC Web server terminated")
}
