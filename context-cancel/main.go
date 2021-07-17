package main

import (
	"context"
	"log"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	go contextCancel(ctx)

	time.Sleep(time.Second*15)
}

func contextCancel(ctx context.Context) {
	select {
	case <- ctx.Done():
		log.Println("errrr")
	default:
		rawFunction()
	}
}

func rawFunction() string {
	time.Sleep(time.Second * 5)

	return "Hey man"
}
