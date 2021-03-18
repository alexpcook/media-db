package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/schema"
)

func TestDelete(tt *testing.T) {
	cfg, err := config.LoadMediaDbConfig()
	if err != nil {
		tt.Fatal(err)
	}

	client, err := NewMediaDbClient(cfg)
	if err != nil {
		tt.Fatal(err)
	}

	movie, err := schema.NewMovie("A Movie Title", "A Movie Director", 2010, "2021-02-16")
	if err != nil {
		tt.Fatal(err)
	}

	err = client.Delete(movie)
	if err != nil {
		tt.Fatal(err)
	}

	music, err := schema.NewMusic("An Album Title", "An Artist", 1980, "2020-03-16")
	if err != nil {
		tt.Fatal(err)
	}

	err = client.Delete(music)
	if err != nil {
		tt.Fatal(err)
	}
}
