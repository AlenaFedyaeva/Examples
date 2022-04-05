package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"greet/greetpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Srv struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *Srv) Greet(ctx context.Context,
	req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	fmt.Println("Greet client:", req, req.Greeting.FirstName)
	firstName := req.GetGreeting().FirstName

	result := "Hello " + firstName
	return &greetpb.GreetResponse{Result: result}, nil
}

func (s *Srv) GreetManyTimes(req *greetpb.GreetRequestManyTimes,
	stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Println("Greet client MT:",
		req, req.GreetingMT.FirstName)
	fn := req.GreetingMT.FirstName
	ln := req.GreetingMT.LastName

	for i := 0; i < 10; i++ {
		str := fmt.Sprintf("%d HelloMT %s %s ", i, fn, ln)
		res := &greetpb.GreetResponseManyTimes{
			Result: str,
		}
		stream.Send(res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil
}

func (s *Srv) GreetLong(stream greetpb.GreetService_GreetLongServer) error {
	fmt.Println("Greet client streaming")
	result := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//we have finished reading the client stream
			return stream.SendAndClose(&greetpb.GreetResponseCliStriaming{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		name := req.GetGreeting().StreamName
		result = result + " " + name + "|"
		fmt.Printf("Recv stream: %s \n", result)
	}

}

//Bi Direct. streaming
func (s *Srv) GreetBiDirect(stream greetpb.GreetService_GreetBiDirectServer) error {
	fmt.Println("Greet bi direct streaming")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading bi direct stream: %v", err)
			return err
		}

		result := "Hello " + req.GetGreetingBd().FirstName + " " +
			req.GetGreetingBd().FirstName

		err = stream.Send(
			&greetpb.GreetResponseBiDirecrt{
				Result: result,
			})

		if err != nil {
			log.Fatalf("err while sending bi direct data: %v", err)
			return err
		}
		fmt.Println(result)

	}
}

func (s *Srv) GreetErr(ctx context.Context,
	req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

	name := req.Greeting.GetFirstName()
	if name == "Eric" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recv forbidden name: %v", name),
		)
	}
	return &greetpb.GreetResponse{
		Result: name,
	}, nil
}

func (s *Srv) GreetDeadline(ctx context.Context,
	req *greetpb.GreetRequestDeadline) (*greetpb.GreetResponseDeadline, error) {

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.Canceled {
			fmt.Println("Client canceled the request")
			return nil,
				status.Error(codes.Canceled,
					"the client canceled the request")
		}
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Greet client:", req, req.GreetingDeadline.FirstName)
	firstName := req.GetGreetingDeadline().FirstName

	result := "Hello " + firstName
	return &greetpb.GreetResponseDeadline{Result: result}, nil
}

func main() {
	fmt.Println("Server greet start...", greetpb.Day)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &Srv{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
