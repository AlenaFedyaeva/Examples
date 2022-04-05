package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	fmt.Println("Client starting...")

	ip := "localhost:50051"
	cc, err := grpc.Dial(ip, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	//doUnary(c)
	//doServerStreaming(c)
	//doClientStreaming(c)
	//doBiDirectStreaming(c)
	//doErrUnary(c)

	//Dealines
	//doUnaryDeadline(c, 5*time.Second)
	doUnaryDeadline(c, 1*time.Millisecond)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("start grpc client...")
	in := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Alena",
			LastName:  "Fd",
		},
	}

	res, err := c.Greet(context.Background(), in)
	if err != nil {
		log.Printf("client created %v", err)
	}

	log.Printf("Resp from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("start grpc client...")
	in := &greetpb.GreetRequestManyTimes{
		GreetingMT: &greetpb.GreetingManyTimes{
			FirstName: "Nella",
			LastName:  "Yam",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), in)
	if err != nil {
		log.Printf("client strean RPC %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("err while reading stream")
		}
		log.Println("RESP from GreatManyTimes msg", msg.String())
	}

}

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting streaming cli RPC...")

	requests := []*greetpb.GreetRequestCliStreaming{
		&greetpb.GreetRequestCliStreaming{
			Greeting: &greetpb.GreetingCliStreaming{
				StreamName: "Name1",
			},
		},
		&greetpb.GreetRequestCliStreaming{
			Greeting: &greetpb.GreetingCliStreaming{
				StreamName: "Name2",
			},
		},
		&greetpb.GreetRequestCliStreaming{
			Greeting: &greetpb.GreetingCliStreaming{
				StreamName: "Name3",
			},
		},
		&greetpb.GreetRequestCliStreaming{
			Greeting: &greetpb.GreetingCliStreaming{
				StreamName: "Name4",
			},
		},
	}

	stream, err := c.GreetLong(context.Background())
	if err != nil {
		log.Fatalf("error while calling ClientStreaming %v", err)
	}

	for _, req := range requests {
		fmt.Println("Send req:", req)
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while  close streaming %v", err)
	}

	fmt.Printf("GreetLong streaming %v\n", res)

}

// Bi Directional streaming
func doBiDirectStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting streaming cli RPC...")

	requests := []*greetpb.GreetRequestBiDirect{
		&greetpb.GreetRequestBiDirect{
			GreetingBd: &greetpb.Greeting{
				FirstName: "Who1",
				LastName:  "what1",
			},
		},
		&greetpb.GreetRequestBiDirect{
			GreetingBd: &greetpb.Greeting{
				FirstName: "Who2",
				LastName:  "what2",
			},
		},
		&greetpb.GreetRequestBiDirect{
			GreetingBd: &greetpb.Greeting{
				FirstName: "Who3",
				LastName:  "what3",
			},
		},
		&greetpb.GreetRequestBiDirect{
			GreetingBd: &greetpb.Greeting{
				FirstName: "Who4",
				LastName:  "what4",
			},
		},
	}

	stream, err := c.GreetBiDirect(context.Background())
	if err != nil {
		log.Fatalf("error while calling Streaming %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Println("Send req:", req)
			stream.Send(req)
			time.Sleep(100 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {

		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("err while rec: %v", err)
				break
			}
			fmt.Printf("Recv Bi Drect: %v \n", res.GetResult())
		}
		close(waitc)

	}()

	// block until everything is done
	<-waitc

}

func doErrUnary(c greetpb.GreetServiceClient) {
	fmt.Println("start grpc client...")
	inValid := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Alice",
			LastName:  "Fd",
		},
	}

	doErrorCall(c, inValid)

	inErr := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Eric",
			LastName:  "Fd",
		},
	}
	doErrorCall(c, inErr)

}

func doErrorCall(c greetpb.GreetServiceClient,
	in *greetpb.GreetRequest) {

	res, err := c.GreetErr(context.Background(),
		in)
	if err != nil {
		log.Printf("client created %v", err)
		respErr, ok := status.FromError(err)
		if ok {

			fmt.Println(respErr.Message())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("You used wrong name Eric", respErr.Message())
			}

		} else {
			log.Fatalf("Big err %v", err)
		}

		log.Fatalf("Other err %v", err)
		return
	}

	fmt.Printf("Resp from ErrGreet: %v", res.GetResult())

}

func doUnaryDeadline(c greetpb.GreetServiceClient, seconds time.Duration) {
	fmt.Println("start grpc client Deadline ... ")
	in := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Alena",
			LastName:  "Fd",
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(),
		1*time.Millisecond)

	defer cancel()

	res, err := c.Greet(ctx, in)
	if err != nil {
		statusErr, ok := status.FromError(err)

		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout was hit. Deadline was exceded")
			} else {
				fmt.Println("Deadline unexpected err: %v ", err)
			}
		} else {
			log.Fatal("err while calling Deadline RPC: %v", err)
		}
		return

	}

	log.Printf("Resp from eadline: %v", res.Result)
}
