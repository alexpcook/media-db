package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
)

func TestView(tt *testing.T) {
	preTestSetup()
	defer postTestTeardown()

	cfg, err := config.LoadMediaDbConfig()
	if err != nil {
		tt.Fatal(err)
	}

	client, err := NewMediaDbClient(cfg)
	if err != nil {
		tt.Fatal(err)
	}

	res, err := client.View("")
	if err != nil {
		tt.Fatal(err)
	}

	if len(res) != 2 {
		tt.Fatalf("expected two entries in response, got %d", len(res))
	}
}
