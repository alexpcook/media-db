package service

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/alexpcook/media-db/config"
)

var (
	testConfigFile *string = flag.String("config", "", "The media db configuration file to use for testing")
)

func preTestSetup() {
	flag.Parse()

	if *testConfigFile == "" {
		log.Fatalf("must specify a media db configuration file for testing, got %q", *testConfigFile)
	}

	err := os.Remove(config.GetDefaultConfigFile())
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
		log.Println(err)
	}

	err = os.Setenv(config.GetOverrideConfigFileEnvVar(), *testConfigFile)
	if err != nil {
		log.Fatal(err)
	}
}
