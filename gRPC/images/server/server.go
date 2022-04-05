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
