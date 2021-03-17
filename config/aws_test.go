package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestGetConfigFilePath(tt *testing.T) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		tt.Fatal(err)
	}

	testCases := []struct {
		dir, file, want string
	}{
		{"dir1", "file1", userHomeDir + "/dir1" + "/file1"},
		{"//dir2", "/file2/", userHomeDir + "/dir2" + "/file2"},
		{DefaultConfigDir, DefaultConfigFile, fmt.Sprintf("%s/%s/%s", userHomeDir, DefaultConfigDir, DefaultConfigFile)},
	}

	for i, test := range testCases {
		tt.Run(fmt.Sprintf("case-%d", i), func(subtt *testing.T) {
			filepath, err := GetConfigFilePath(test.dir, test.file)
			if err != nil {
				subtt.Fatal(err)
			}

			if test.want != filepath {
				subtt.Fatalf("want %q, got %q", test.want, filepath)
			}
		})
	}
}

func invalidFilePathTestName() string {
	return "invalid-filepath"
}

func TestLoadMediaDbConfig(tt *testing.T) {
	testCases := []struct {
		name    string
		json    []byte
		config  *MediaDbConfig
		isError bool
	}{
		{
			"basic",
			[]byte(`{"profile": "test-profile", "region": "us-west-1", "bucket": "test-bucket"}`),
			&MediaDbConfig{AWSProfile: "test-profile", AWSRegion: "us-west-1", S3Bucket: "test-bucket"},
			false,
		},
		{
			invalidFilePathTestName(),
			[]byte(`{"profile": "test-profile", "region": "us-west-1", "bucket": "test-bucket"}`),
			nil,
			true,
		},
		{
			"invalid-json",
			[]byte(`{"profile: "test-profile", "region": "us-west-1", "bucket": "test-bucket"}`),
			nil,
			true,
		},
		{
			"null-profile",
			[]byte(`{"profile": "    ", "region": "us-west-1", "bucket": "test-bucket"}`),
			nil,
			true,
		},
		{
			"null-region",
			[]byte(`{"profile": "test-profile", "region": "   ", "bucket": "test-bucket"}`),
			nil,
			true,
		},
		{
			"null-bucket",
			[]byte(`{"profile": "test-profile", "region": "us-west-1", "bucket": ""}`),
			nil,
			true,
		},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			tempConfigFile, err := os.CreateTemp("", "aws_test")
			if err != nil {
				subtt.Fatal(err)
			}
			defer func() {
				err = os.Remove(tempConfigFile.Name())
				if err != nil {
					subtt.Fatal(err)
				}
			}()

			_, err = tempConfigFile.Write(test.json)
			if err != nil {
				subtt.Fatal(err)
			}

			filepath := tempConfigFile.Name()
			if test.name == invalidFilePathTestName() {
				filepath = "an invalid filepath"
			}

			got, err := LoadMediaDbConfig(filepath)
			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if !reflect.DeepEqual(test.config, got) {
				subtt.Fatalf("want %v, got %v", test.config, got)
			}
		})
	}
}

func TestLoadAWSConfig(tt *testing.T) {
	testCases := []struct {
		name    string
		db      *MediaDbConfig
		isError bool
	}{
		{
			"basic",
			&MediaDbConfig{AWSProfile: "test-profile", AWSRegion: "us-west-1", S3Bucket: "test-bucket"},
			false,
		},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			cfg, err := test.db.LoadAWSConfig()

			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if test.db.AWSRegion != cfg.Region {
				subtt.Fatalf("error with region: want %s, got %s", test.db.AWSRegion, cfg.Region)
			}
		})
	}
}
