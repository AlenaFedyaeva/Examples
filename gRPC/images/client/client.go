package main

import (
	"context"
	"fmt"
	"imagegrpc/client/consumer"
	"imagegrpc/internal/use"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

var workerPoolSize = 10

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	use.SetLogLevel()
	log.Info().Msg("Client starting...")

	ip := use.GetIpAddress()
	cc, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("could not connect: %v", err)
	}
	defer cc.Close()

	client := consumer.NewClient(cc)

	wg := &sync.WaitGroup{}
	wg.Add(workerPoolSize)
	// Start consumer with cancellation context passed
	go client.StartConsumer(ctx)

	//Start [workerPoolSize] workers
	for i := 0; i < workerPoolSize; i++ {
		go client.WorkerFunc(wg, i)
	}

	go client.WorkerAddOpt(ctx)

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	//  go startConsumer(ctx)
	<-termChan // Blocks here until either SIGINT or SIGTERM is received.
	cancelFunc()
	wg.Wait() // program will wait here until all worker goroutines have reported that they're done
	fmt.Println("Workers done. Bye")
}
