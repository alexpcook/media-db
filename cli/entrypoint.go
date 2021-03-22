package cli

import (
	"log"
	"os"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/service"
)

var (
	MediaDbConfig *config.MediaDbConfig
	MediaDbClient *service.MediaDbClient

	StdoutLogger *log.Logger = log.New(os.Stdout, "", 0)
	StderrLogger *log.Logger = log.New(os.Stderr, "", 0)
)

func init() {
	MediaDbConfig, err := config.LoadMediaDbConfig()
	if err != nil {
		StderrLogger.Fatal(err)
	}

	MediaDbClient, err = service.NewMediaDbClient(MediaDbConfig)
	if err != nil {
		StderrLogger.Fatal(err)
	}
}

func Execute() {
	if len(os.Args) < 2 {
		StderrLogger.Fatal(GetCLIHelpText())
	}

	cmd, err := NewMediaDbCommand(os.Args[1:])
	if err != nil {
		StderrLogger.Fatal(err)
	}

	err = cmd.Run()
	if err != nil {
		StderrLogger.Fatal(err)
	}
}
