package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
	"github.com/alexpcook/media-db-console/schema"
)

func TestCreate(tt *testing.T) {
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

	err = client.Create(movie)
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = client.Delete(movie.ID, *movie)
		if err != nil {
			tt.Fatal(err)
		}
	}()

	music, err := schema.NewMusic("An Album Title", "An Artist", 1980, "2020-03-16")
	if err != nil {
		tt.Fatal(err)
	}

	err = client.Create(music)
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = client.Delete(music.ID, *music)
		if err != nil {
			tt.Fatal(err)
		}
	}()

	// Mock a failed communication with the S3 bucket.
	originalS3Bucket := client.s3Bucket
	client.s3Bucket = "this-is-an-invalid-bucket-name"
	err = client.Create(music)
	if err == nil {
		tt.Fatal("want error, got nil")
	}
	defer func() {
		client.s3Bucket = originalS3Bucket
	}()
}
