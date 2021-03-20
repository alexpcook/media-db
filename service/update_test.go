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

	// Ensure that the database is empty.
	entries, err := client.Read("")
	if err != nil {
		tt.Fatal(err)
	}
	if len(entries) != 0 {
		tt.Fatalf("want 0 entries in database, got %d", len(entries))
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

	// Validate that one film is now in the database.
	entries, err = client.Read("")
	if err != nil {
		tt.Fatal(err)
	}
	if len(entries) != 1 {
		tt.Fatalf("want 1 entry in database, got %d", len(entries))
	}

	for _, entry := range entries {
		switch entry.(type) {
		case schema.Movie:
			break
		default:
			tt.Fatalf("expected movie type, got %T", entry)
		}
	}
}
