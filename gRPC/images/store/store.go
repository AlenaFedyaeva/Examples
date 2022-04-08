package store

import (
	"bytes"
	"fmt"
	"time"

	"os"
	"sync"

	"github.com/rs/zerolog/log"

	"github.com/google/uuid"
)

type DiskImageStore struct {
	mutex       sync.RWMutex
	imageFolder string
	images      map[string]*ImageInfo
}

type ImageInfo struct {
	Path       string
	CreateTime time.Time
	UpdateTime time.Time
}

func MakeDirectoryIfNotExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModeDir|0755)
	}
	log.Info().Str("folder name", path).Msg("folder created")

	return nil
}

func NewDiskImageStore(imageFolder string) (*DiskImageStore, error) {
	err := MakeDirectoryIfNotExists(imageFolder)
	if err != nil {
		return nil, err

	}

	return &DiskImageStore{
		imageFolder: imageFolder,
		images:      make(map[string]*ImageInfo),
	}, nil
}

func (store *DiskImageStore) Save(imageName string,
	imageData bytes.Buffer) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageID, imageName)

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.images[imageID.String()] = &ImageInfo{
		Path:       imagePath,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}

	return imageID.String(), nil
}
