package main

import (
	"backend/api/playground"
	"backend/config"
	"backend/internal/db/redis_service"
	"backend/internal/endpoints"
	playgroundsvc "backend/internal/services/playground"
	"backend/internal/transports"
	"backend/pkg/utilities"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/configor"
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	
	cfg := loadConfig()
	
	gRPCTransport := transports.BaseGRPCTransport(cfg) // transport layer
	gRPCWebTransport := transports.BaseGRPCWebTransport(gRPCTransport.Server(), cfg)
	
	redisClient := GetRedisCli()
	
	rds := redis_service.New(redisClient) // db repository layer
	
	pgService, _ := playgroundsvc.New(rds)                                      // service layer
	pgEndpoints := endpoints.NewPlaygroundEndpoint(pgService)                // endpoints layer
	playground.RegisterPlaygroundServer(gRPCTransport.Server(), pgEndpoints) // registration
	
	// run transports
	gRPCTransport.Run()
	gRPCWebTransport.Run()
	
	// infinite wait
	select {}
}

func GetRedisCli() *redis.Client {
	// Redis Connection
	cli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", "0.0.0.0", "6379"),
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
