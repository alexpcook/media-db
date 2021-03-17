package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/schema"
)

func handleError(err error, tt *testing.T) {
	if err != nil {
		tt.Fatal(err)
	}
}

func TestAdd(tt *testing.T) {
	if *testConfigFilepath == "" {
		tt.Fatalf("a media db config file for testing must be supplied, got %q", *testConfigFilepath)
	}

	cfg, err := config.LoadMediaDbConfig(*testConfigFilepath)
	handleError(err, tt)

	client, err := NewMediaDbClient(cfg)
	handleError(err, tt)

	movie, err := schema.NewMovie("A Movie Title", "A Movie Director", 2010, "2021-02-16")
	handleError(err, tt)

	err = client.Add(movie)
	handleError(err, tt)

	music, err := schema.NewMusic("An Album Title", "An Artist", 1980, "2020-03-16")
	handleError(err, tt)

	err = client.Add(music)
	handleError(err, tt)
}
