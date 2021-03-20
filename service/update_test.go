package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/schema"
)

func TestUpdate(tt *testing.T) {
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

	// Try to update the movie before it exists to force an error.
	err = client.Update(movie.ID, *movie)
	if err == nil {
		tt.Fatal("want error, got nil")
	}

	err = client.Create(movie)
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = client.Delete(movie)
		if err != nil {
			tt.Fatal(err)
		}
	}()

	// Update the year and sync it to the database.
	movie.YearMade = 2020
	err = client.Update(movie.ID, *movie)
	if err != nil {
		tt.Fatal(err)
	}
}
