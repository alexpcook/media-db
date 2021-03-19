package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/schema"
)

func TestDelete(tt *testing.T) {
	// Simulate a failed delete.
	// Successful deletes are tested elsewhere in the package.
	cfg, err := config.LoadMediaDbConfig()
	if err != nil {
		tt.Fatal(err)
	}

	client, err := NewMediaDbClient(cfg)
	if err != nil {
		tt.Fatal(err)
	}

	movie, err := schema.NewMovie("Some Movie Title", "Some Movie Director", 2010, "2021-02-16")
	if err != nil {
		tt.Fatal(err)
	}

	client.s3Bucket = "this-is-an-invalid-bucket-name"
	err = client.Delete(movie)
	if err == nil {
		tt.Fatal("want error, got nil")
	}
}
