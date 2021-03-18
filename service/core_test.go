package service

import (
	"errors"
	"flag"
	"log"
	"os"
	"testing"

	"github.com/alexpcook/media-db-console/config"
)

var (
	testConfigFile string
)

func init() {
	flag.StringVar(&testConfigFile, "conf", "", "The media db config file to use for testing")

	err := os.Unsetenv(config.GetOverrideConfigFileEnvVar())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(config.GetDefaultConfigFile())
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
		log.Println(err)
	}
}

func TestNewMediaDbClient(tt *testing.T) {
	flag.Parse()

	if testConfigFile == "" {
		tt.Fatalf("a media db config file for testing must be supplied, got %q", testConfigFile)
	}

	err := os.Setenv(config.GetOverrideConfigFileEnvVar(), testConfigFile)
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = os.Unsetenv(config.GetOverrideConfigFileEnvVar())
		if err != nil {
			tt.Fatal(err)
		}
	}()

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
