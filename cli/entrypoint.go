package cli

import (
	"log"
	"os"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/service"
)

const (
	Create string = "create"
	Read   string = "read"
	Update string = "update"
	Delete string = "delete"
)

var (
	MediaDbConfig *config.MediaDbConfig
	MediaDbClient *service.MediaDbClient
)

func init() {
	MediaDbConfig, err := config.LoadMediaDbConfig()
	if err != nil {
		log.Fatal(err)
	}

	MediaDbClient, err = service.NewMediaDbClient(MediaDbConfig)
	if err != nil {
		log.Fatal(err)
	}
}

func Execute() {
	if len(os.Args) < 2 {
		log.Fatal("no command given")
	}

	switch subCommand := os.Args[1]; subCommand {
	case Create, Read, Update, Delete:
		log.Println(subCommand)
	default:
		log.Fatalf("%s is not a valid command", subCommand)
	}
}
