package main

import (
	"backend/config"
	"backend/internal/core/application"
	"backend/internal/outside/adapter/driven/persistance"
	"backend/internal/outside/adapter/driven/sandbox"
	"backend/internal/outside/adapter/driving/grpchandler"
	"backend/internal/pkg/logger"
	"backend/internal/pkg/middlewares"
	"backend/internal/pkg/transports/grpc"
	"backend/internal/pkg/transports/grpcgateway"
	"backend/internal/pkg/transports/grpcweb"
	"backend/internal/pkg/transports/http"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

var GitCommit string

func main() {
	logger.New()
	logger.Log.Println("Git tag ver:", GitCommit)

	_ = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	runtime.GOMAXPROCS(1)

	// =========================================================================
	// 3rd Party Services Layer
	// =========================================================================

	// Redis Connection
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Cfg.RedisHost, config.Cfg.RedisPort),
	})

	ping, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ping)

	// =========================================================================
	// Adapter Layer
	// =========================================================================

	codeRepository := persistance.NewCodeDBRedis(redisClient)
	sandboxRepository := sandbox.NewCMDBasedSandbox()

	// =========================================================================
	// Middleware Layer
	// =========================================================================

	mws := middlewares.NewMiddleware()

	// =========================================================================
	// Transport Layer
	// =========================================================================

	grpcTransport, err := grpc.New(mws)
	if err != nil {
		log.Fatal(err)
	}

	// grpc gateway REST transport
	grpcGatewayTransport := grpcgateway.NewGRPCGateway(mws)

	// grpc web transport
	grpcWebTransport := grpcweb.NewGRPCWeb(grpcTransport.Server())

	// http transport for web
	httpTransport := http.New(mws)

	// =========================================================================
	// Services Layer
	// =========================================================================

	codeService := application.NewCodeService(codeRepository, sandboxRepository)

	// =========================================================================
	// Ports Layer
	// =========================================================================

	codeGRPC := grpchandler.NewCodeGRPC(codeService)
	codeGRPC.Wire(grpcTransport.Server(), grpcGatewayTransport.Server())

	// =========================================================================
	// Transport Runner Layer
	// =========================================================================

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGSTOP)

	// run grpc
	grpcTransport.Run()

	// run grpc gateway
	grpcGatewayTransport.Run()

	// run grpc web
	grpcWebTransport.Run()

	// run http transport
	httpTransport.Run()

	// =========================================================================
	// Transport Halting Layer
	// =========================================================================

	shutdownCtx := context.Background()

	<-sigChan
	grpcTransport.Shutdown(shutdownCtx)
	grpcGatewayTransport.Shutdown(shutdownCtx)
	grpcWebTransport.Shutdown(shutdownCtx)
	httpTransport.Shutdown(shutdownCtx)
}
