package cli

import (
	"fmt"
	"log"
	"os"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/service"
)

var (
	// MediaDbConfig is the media db configuration for the CLI.
	MediaDbConfig *config.MediaDbConfig
	// MediaDbClient is the media db service client for the CLI.
	MediaDbClient *service.MediaDbClient

	// StdoutLogger is the standard output logger for the CLI.
	StdoutLogger *log.Logger = log.New(os.Stdout, "", 0)
	// StderrLogger is the standard error logger for the CLI.
	StderrLogger *log.Logger = log.New(os.Stderr, "", 0)
)

// InitDb loads the media database configuration and initializes
// a service client to communicate with the database. It will return
// a non-nil error if any of these steps fail.
func InitDb() {
	MediaDbConfig, err := config.LoadMediaDbConfig()
	if err != nil {
		StderrLogger.Fatal(fmt.Sprintf(`%s
		
run 'mdb setup' to fix the configuration issue`, err.Error()))
	}

	MediaDbClient, err = service.NewMediaDbClient(MediaDbConfig)
	if err != nil {
		StderrLogger.Fatal(fmt.Sprintf(`%s
		
run 'mdb setup' to fix the configuration issue`, err.Error()))
	}
}

// Execute is the main entrypoint for callers.
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
