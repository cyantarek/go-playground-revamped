package main

import (
	"backend/config"
	"backend/internal/db/redis_service"
	"backend/internal/endpoints/rpc"
	playgroundsvc "backend/internal/services/playground"
	"backend/internal/transports"
	"backend/pkg/utilities"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/configor"
	"google.golang.org/grpc"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	cfg := loadConfig()

	// transports layers
	gRPCTransport := transports.BaseGRPCTransport(cfg)
	gRPCWebTransport := transports.BaseGRPCWebTransport(gRPCTransport.Server(), cfg)
	httpTransport := transports.BaseHttpTransport(cfg)

	// db repository layer
	redisClient := GetRedisCli()
	rds := redis_service.New(redisClient)

	// services layer
	pgService, _ := playgroundsvc.New(rds)

	// endpoints layer
	pgEndpoints := rpc.NewPlaygroundEndpoints(pgService)

	// registration
	gRPCTransport.Register(pgEndpoints)

	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := httpTransport.Register(fmt.Sprintf("%s:%s", cfg.Server.Grpc.Host, cfg.Server.Grpc.Port), opts)
	if err != nil {
		log.Fatal(err)
	}

	// run transports in separate goroutines
	gRPCTransport.Run()
	gRPCWebTransport.Run()
	httpTransport.Run()

	// infinite wait
	select {}
}

func GetRedisCli() *redis.Client {
	// Redis Connection
	cli := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "0.0.0.0", "6379"),
	})

	ping, err := cli.Ping().Result()
	if err != nil {
		fmt.Println(fmt.Sprintf("Redis Service is Offline :  %s \n", err.Error()))
	}
	fmt.Println(ping)

	return cli
}

func loadConfig() *config.Config {
	var cfg config.Config
	env := utilities.GetEnv("ENV", "dev")

	switch env {
	case "dev":
		err := configor.Load(&cfg, "config/config.dev.yml")
		if err != nil {
			log.Fatal(err)
		}
	case "prod":
		err := configor.Load(&cfg, "config/config.prod.yml")
		if err != nil {
			log.Fatal(err)
		}
	}

	return &cfg
}
