package service

import (
	"testing"

	"github.com/alexpcook/media-db-console/config"
)

func TestView(tt *testing.T) {
	if *testConfigFilepath == "" {
		tt.Fatalf("a media db config file for testing must be supplied, got %q", *testConfigFilepath)
	}

	cfg, err := config.LoadMediaDbConfig()
	handleError(err, tt)

	client, err := NewMediaDbClient(cfg)
	handleError(err, tt)

	res, err := client.View("")
	handleError(err, tt)

	if len(res) != 2 {
		tt.Fatalf("expected two entries in response, got %d", len(res))
	}
}
