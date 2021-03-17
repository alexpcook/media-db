package service

import (
	"os"
	"testing"

	"github.com/alexpcook/media-db-console/config"
)

func getLocalEnvVar(s string, tt *testing.T) string {
	val, isSet := os.LookupEnv(s)
	if !isSet {
		tt.Fatalf("env variable %q not set", s)
	}

	return val
}

func TestNewMediaDbClient(tt *testing.T) {
	testCases := []struct {
		name          string
		mediaDbConfig *config.MediaDbConfig
		isError       bool
	}{
		{
			"basic",
			&config.MediaDbConfig{
				AWSProfile: getLocalEnvVar("MEDIA_DB_TEST_AWS_PROFILE", tt),
				AWSRegion:  getLocalEnvVar("MEDIA_DB_TEST_AWS_REGION", tt),
				S3Bucket:   getLocalEnvVar("MEDIA_DB_TEST_S3_BUCKET", tt),
			},
			false,
		},
		{
			"invalid-profile",
			&config.MediaDbConfig{
				AWSProfile: "not a profile",
				AWSRegion:  getLocalEnvVar("MEDIA_DB_TEST_AWS_REGION", tt),
				S3Bucket:   getLocalEnvVar("MEDIA_DB_TEST_S3_BUCKET", tt),
			},
			true,
		},
		{
			"invalid-region",
			&config.MediaDbConfig{
				AWSProfile: getLocalEnvVar("MEDIA_DB_TEST_AWS_PROFILE", tt),
				AWSRegion:  "not a region",
				S3Bucket:   getLocalEnvVar("MEDIA_DB_TEST_S3_BUCKET", tt),
			},
			true,
		},
		{
			"invalid-bucket",
			&config.MediaDbConfig{
				AWSProfile: getLocalEnvVar("MEDIA_DB_TEST_AWS_PROFILE", tt),
				AWSRegion:  getLocalEnvVar("MEDIA_DB_TEST_AWS_REGION", tt),
				S3Bucket:   "not a bucket",
			},
			true,
		},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			client, err := NewMediaDbClient(test.mediaDbConfig)

			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if test.mediaDbConfig.S3Bucket != client.s3Bucket {
				subtt.Fatalf("wrong s3 bucket name: want %s, got %s", test.mediaDbConfig.S3Bucket, client.s3Bucket)
			}
		})
	}
}
