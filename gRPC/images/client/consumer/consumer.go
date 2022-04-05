package consumer

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"imagegrpc/imagepb"
	"imagegrpc/internal/use"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

const workerPoolSize = 10

type Options struct {
	opt string
	val string
}

func newOpt() (Options, error) {
	fname := ""
	opt, err := use.GetOption()
	if err != nil {
		return Options{}, err
	}

	switch opt {
	case use.OptLs:
	case use.OptUpload:
		fname, err = use.GetUploadFileName()
		if err != nil {
			return Options{}, err
		}
	default:
		return Options{}, fmt.Errorf("Wrong option %s", opt)
	}

	return Options{
		opt: opt,
		val: fname,
	}, nil
}

type Client struct {
	ingestChanUpload chan string
	ingestChanLs     chan string
	jobsChanUpload   chan string
	jobsChanLs       chan string
	client           imagepb.ImageServiceClient
}

func NewClient(cc *grpc.ClientConn) Client {
	return Client{
		ingestChanUpload: make(chan string, 1),
		ingestChanLs:     make(chan string, 1),
		jobsChanUpload:   make(chan string, workerPoolSize),
		jobsChanLs:       make(chan string, 100),
		client:           imagepb.NewImageServiceClient(cc),
	}
}


func (c Client) WorkerAddOpt(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChanUpload)
			close(c.jobsChanLs)
			fmt.Println("Consumer closed jobsChan")
			return
		default:
			options, err := newOpt()
			if err != nil {
				log.Error().Err(err).Send()
			}
			switch options.opt {
			case use.OptUpload:
				c.ingestChanUpload <- options.val
				// case use.OptLs:
				// 	c.ingestChanUpload <- options.opt
			}
		}
	}
}


func (c Client) WorkerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	log.Info().Msgf("Worker %d starting\n", index)
	for fname := range c.jobsChanUpload {
		log.Info().Msgf("Worker %d started job %d\n", index, fname)
		c.uploadImage(fname)
		
		log.Info().Msgf("Worker %d finished processing job %d\n", index, fname)
	}
	fmt.Printf("Worker %d interrupted\n", index)
}


func (c Client) StartConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.ingestChanUpload:
			c.jobsChanUpload <- job
		case <-ctx.Done():
			fmt.Println("Consumer received cancellation signal, closing jobsChan!")
			close(c.jobsChanUpload)
			close(c.jobsChanLs)
			fmt.Println("Consumer closed jobsChan")
			return
		}
	}
}

//  upload image on server
//  Users/alfedyaeva/workspace/tmp/images/client/img/big_flower.png
func (c Client) uploadImageFromStdin() error {
	imagePath, err := use.GetUploadFileName()
	if err != nil {
		return err
	}

	return c.uploadImage(imagePath)
}

func (c Client) uploadImage(imagePath string) error {
	log.Info().Str("path", imagePath).Msg("upload image")

	if _, err := os.Stat(imagePath); errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("image do not exist", imagePath)
	}

	// 1) open stream
	file, err := os.Open(imagePath)
	if err != nil {
		return fmt.Errorf("cannot open image file: %v ", err)
	}
	defer file.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Second)
	defer cancel()

	stream, err := c.client.ImageLoadOnServer(ctx)
	if err != nil {
		return fmt.Errorf("cannot upload image: %v", err)
	}

	//  2) send file name
	req := &imagepb.UploadImageRequest{
		Data: &imagepb.UploadImageRequest_Info{
			Info: &imagepb.ImageInfo{
				ImageName: filepath.Ext(imagePath),
			},
		},
	}

	err = stream.Send(req)
	if err != nil {
		return fmt.Errorf("cannot send image info to server: %v", err)
	}

	// 3) send file body
	reader := bufio.NewReader(file)
	buffer := make([]byte, 20)

	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			return fmt.Errorf("cannot read chunk to buffer: %v", err)
		}

		req := &imagepb.UploadImageRequest{
			Data: &imagepb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		err = stream.Send(req)
		if err != nil {
			return fmt.Errorf("cannot send chunk to server: %v ", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return fmt.Errorf("cannot receive response: %v", err)
	}

	log.Info().Str("image", imagePath).Str("id", res.GetId()).Uint32("size", res.GetSize()).Msg("saved succesful")

	return nil
}
