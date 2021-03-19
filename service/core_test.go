package service

import (
	"os"
	"testing"

	"github.com/alexpcook/media-db-console/config"
)

func TestMain(m *testing.M) {
	preTestSetup()
	os.Exit(m.Run())
}

func TestNewMediaDbClient(tt *testing.T) {
	cfg, err := config.LoadMediaDbConfig()
	if err != nil {
		tt.Fatal(err)
	}

	testCases := []struct {
		name    string
		isError bool
	}{
		{"basic", false},
		{"invalid-profile", true},
		{"invalid-region", true},
		{"invalid-bucket", true},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			switch test.name {
			case "invalid-profile":
				cfg.AWSProfile = "not-an-aws-profile"
			case "invalid-region":
				cfg.AWSRegion = "not-an-aws-region"
			case "invalid-bucket":
				cfg.S3Bucket = "not-an-s3-bucket-ajfkldfjadslfdsflkdsajfdsk"
			}

			client, err := NewMediaDbClient(cfg)

			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if cfg.S3Bucket != client.s3Bucket {
				subtt.Fatalf("wrong s3 bucket name: want %s, got %s", cfg.S3Bucket, client.s3Bucket)
			}
		})
	}
}
