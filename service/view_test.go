package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/schema"
)

func TestView(tt *testing.T) {
	cfg, err := config.LoadMediaDbConfig()
	if err != nil {
		tt.Fatal(err)
	}

	client, err := NewMediaDbClient(cfg)
	if err != nil {
		tt.Fatal(err)
	}

	movie, err := schema.NewMovie("Another Title", "Another Director", 1965, "2019-01-13")
	if err != nil {
		tt.Fatal(err)
	}

	err = client.Add(movie)
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = client.Delete(movie)
		if err != nil {
			tt.Fatal(err)
		}
	}()

	res, err := client.View("")
	if err != nil {
		tt.Fatal(err)
	}

	if len(res) != 1 {
		tt.Fatalf("expected one entry in response, got %d", len(res))
	}

	// Test filtering everything in the bucket (since there's only one movie).
	res, err = client.View(schema.GetMediaBaseKey() + "/" + schema.GetMediaTypeKey(schema.Music{}))
	if err != nil {
		tt.Fatal(err)
	}

	if len(res) != 0 {
		tt.Fatalf("expected zero entries in response, got %d", len(res))
	}

	music, err := schema.NewMusic("Another Album", "Another Artist", 2005, "2015-10-31")
	if err != nil {
		tt.Fatal(err)
	}

	err = client.Add(music)
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = client.Delete(music)
		if err != nil {
			tt.Fatal(err)
		}
	}()

	// There should now be one piece of music in the bucket.
	res, err = client.View(schema.GetMediaBaseKey() + "/" + schema.GetMediaTypeKey(schema.Music{}))
	if err != nil {
		tt.Fatal(err)
	}

	if len(res) != 1 {
		tt.Fatalf("expected one entry in response, got %d", len(res))
	}
}
