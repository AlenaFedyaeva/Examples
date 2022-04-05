package main

import (
	"bytes"
	"imagegrpc/imagepb"
	"imagegrpc/internal/use"
	"imagegrpc/store"
	"io"
	"net"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// var once sync.Once

func main() {
	use.SetLogLevel()
	log.Info().Msg(("Server images start..."))

	ip := use.GetIpAddress()
	lis, err := net.Listen("tcp", ip)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	store, err := store.NewDiskImageStore("img")
	if err != nil {
		log.Fatal().Msgf("failed to create storage: %v", err)
	}

	serve := &Srv{
		Store:                           store,
		UnimplementedImageServiceServer: imagepb.UnimplementedImageServiceServer{},
	}

	imagepb.RegisterImageServiceServer(s, serve)

	if err := s.Serve(lis); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}

}

type Srv struct {
	Store ImageStore
	imagepb.UnimplementedImageServiceServer
}

type ImageStore interface {
	Save(imageName string, imageData bytes.Buffer) (string, error)
}

func (s *Srv) ImageLoadOnServer(stream imagepb.ImageService_ImageLoadOnServerServer) error {
	req, err := stream.Recv()
	if err != nil {
		return use.LogError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	fname := req.GetInfo().GetImageName()
	log.Info().Str("receive an upload-image request for file", fname).Send()

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		log.Info().Msg("waiting more data...")

		req, err = stream.Recv()
		if err == io.EOF {
			log.Info().Msg("we have finished reading the client stream")
			break
		}

		if err != nil {
			return status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err)
		}

		chunk := req.GetChunkData()
		size := len(chunk)

		log.Info().Msgf("received a chunk with size: %d", size)

		imageSize += size
		if imageSize > use.MaxImageSize {
			return status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, use.MaxImageSize)
		}
		_, err = imageData.Write(chunk)
		if err != nil {
			return status.Errorf(codes.Internal, "cannot write chunk data: %v", err)
		}

	}

	imageID, err := s.Store.Save(fname, imageData)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot save image to the store: %v", err)
	}

	res := &imagepb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return status.Errorf(codes.Unknown, "cannot send response: %v", err)
	}

	log.Info().Str("image", fname).Str("id", imageID).Int("size", imageSize).Msg("saved succesful")
	return nil
}

// func (s *Srv) ImageGetOnClient(imagepb.ImageService_ImageGetOnClientServer) error {
// 	return nil
// }

// func (s *Srv) Greet(ctx context.Context,
// 	req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

// 	fmt.Println("Greet client:", req, req.Greeting.FirstName)
// 	firstName := req.GetGreeting().FirstName

// 	result := "Hello " + firstName
// 	return &greetpb.GreetResponse{Result: result}, nil
// }

// func (s *Srv) GreetManyTimes(req *greetpb.GreetRequestManyTimes,
// 	stream greetpb.GreetService_GreetManyTimesServer) error {
// 	fmt.Println("Greet client MT:",
// 		req, req.GreetingMT.FirstName)
// 	fn := req.GreetingMT.FirstName
// 	ln := req.GreetingMT.LastName

// 	for i := 0; i < 10; i++ {
// 		str := fmt.Sprintf("%d HelloMT %s %s ", i, fn, ln)
// 		res := &greetpb.GreetResponseManyTimes{
// 			Result: str,
// 		}
// 		stream.Send(res)

// 		time.Sleep(1000 * time.Millisecond)
// 	}

// 	return nil
// }

// func (s *Srv) GreetLong(stream greetpb.GreetService_GreetLongServer) error {
// 	fmt.Println("Greet client streaming")
// 	result := ""

// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			//we have finished reading the client stream
// 			return stream.SendAndClose(&greetpb.GreetResponseCliStriaming{
// 				Result: result,
// 			})
// 		}
// 		if err != nil {
// 			log.Fatalf("Error while reading client stream: %v", err)
// 		}
// 		name := req.GetGreeting().StreamName
// 		result = result + " " + name + "|"
// 		fmt.Printf("Recv stream: %s \n", result)
// 	}

// }

//Bi Direct. streaming
// func (s *Srv) GreetBiDirect(stream greetpb.GreetService_GreetBiDirectServer) error {
// 	fmt.Println("Greet bi direct streaming")

// 	for {
// 		req, err := stream.Recv()
// 		if err == io.EOF {
// 			return nil
// 		}
// 		if err != nil {
// 			log.Fatalf("Error while reading bi direct stream: %v", err)
// 			return err
// 		}

// 		result := "Hello " + req.GetGreetingBd().FirstName + " " +
// 			req.GetGreetingBd().FirstName

// 		err = stream.Send(
// 			&greetpb.GreetResponseBiDirecrt{
// 				Result: result,
// 			})

// 		if err != nil {
// 			log.Fatalf("err while sending bi direct data: %v", err)
// 			return err
// 		}
// 		fmt.Println(result)

// 	}
// }

// func (s *Srv) GreetErr(ctx context.Context,
// 	req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {

// 	name := req.Greeting.GetFirstName()
// 	if name == "Eric" {
// 		return nil, status.Errorf(
// 			codes.InvalidArgument,
// 			fmt.Sprintf("Recv forbidden name: %v", name),
// 		)
// 	}
// 	return &greetpb.GreetResponse{
// 		Result: name,
// 	}, nil
// }

// func (s *Srv) GreetDeadline(ctx context.Context,
// 	req *greetpb.GreetRequestDeadline) (*greetpb.GreetResponseDeadline, error) {

// 	for i := 0; i < 3; i++ {
// 		if ctx.Err() == context.Canceled {
// 			fmt.Println("Client canceled the request")
// 			return nil,
// 				status.Error(codes.Canceled,
// 					"the client canceled the request")
// 		}
// 		time.Sleep(2 * time.Second)
// 	}

// 	fmt.Println("Greet client:", req, req.GreetingDeadline.FirstName)
// 	firstName := req.GetGreetingDeadline().FirstName

// 	result := "Hello " + firstName
// 	return &greetpb.GreetResponseDeadline{Result: result}, nil
// }
