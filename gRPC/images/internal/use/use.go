package use

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var OptLs string = "ls"
var OptUpload string = "upload"

// image size 1 MB = 2^20 byte
const MaxImageSize = 1 << 20

func GetIpAddress() string {
	serverAddress := flag.String("address", "0.0.0.0:50052", "ip address")
	flag.Parse()
	return *serverAddress
}

func SetLogLevel() {
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

func LogError(err error) error {
	if err != nil {
		log.Error().Err(err).Send()
	}
	return err
}

func GetUploadFileName() (string, error) {
	readerStd := bufio.NewReader(os.Stdin)
	fmt.Print("Enter file name: \n")
	line, err := readerStd.ReadString('\n')

	if err != nil {
		return "", err
	}

	imagePath := strings.TrimSuffix(line, "\n")
	log.Info().Str("path", imagePath).Msg("upload image")

	return imagePath, err
}

func GetOption() (string, error) {
	readerStd := bufio.NewReader(os.Stdin)
	fmt.Printf("Options: %s,%s\n", OptLs, OptUpload)
	fmt.Print("Enter option: \n")
	line, err := readerStd.ReadString('\n')

	if err != nil {
		return "", err
	}

	option := strings.TrimSuffix(line, "\n")
	log.Info().Str("option", option).Msg("select option")

	return option, err
}
