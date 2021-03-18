package config

import (
	"errors"
	"log"
	"os"
	"reflect"
	"testing"
)

func preTestSetup() {
	err := os.Unsetenv(GetOverrideConfigFileEnvVar())
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(GetDefaultConfigFile())
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
		log.Println(err)
	}
}

func invalidFilePathTestName() string {
	return "invalid-filepath"
}

func TestLoadMediaDbConfig(tt *testing.T) {
	preTestSetup()

	defaultConfigFile, err := os.Create(GetDefaultConfigFile())
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = os.Remove(defaultConfigFile.Name())
		if err != nil {
			tt.Fatal(err)
		}
	}()

	testCases := []struct {
		name    string
		json    []byte
		want    *MediaDbConfig
		isError bool
	}{
		{
			"basic",
			[]byte(`{"profile": "test-profile", "region": "us-west-1", "bucket": "test-bucket"}`),
			&MediaDbConfig{AWSProfile: "test-profile", AWSRegion: "us-west-1", S3Bucket: "test-bucket"},
			false,
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
			_, err = defaultConfigFile.Write(test.json)
			if err != nil {
				subtt.Fatal(err)
			}

			got, err := LoadMediaDbConfig()
			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if !reflect.DeepEqual(test.want, got) {
				subtt.Fatalf("want %v, got %v", test.want, got)
			}
		})
	}
}

func TestLoadMediaDbConfigOverride(tt *testing.T) {
	preTestSetup()

	testCases := []struct {
		name    string
		json    []byte
		want    *MediaDbConfig
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
			overrideConfigFile, err := os.CreateTemp("", "aws_test_override")
			if err != nil {
				subtt.Fatal(err)
			}
			defer func() {
				err = os.Remove(overrideConfigFile.Name())
				if err != nil {
					subtt.Fatal(err)
				}
			}()

			err = os.Setenv(GetOverrideConfigFileEnvVar(), overrideConfigFile.Name())
			if test.name == invalidFilePathTestName() {
				err = os.Setenv(GetOverrideConfigFileEnvVar(), "/an/invalid/config/file/path")
			}
			if err != nil {
				subtt.Fatal(err)
			}
			defer func() {
				err = os.Unsetenv(GetOverrideConfigFileEnvVar())
				if err != nil {
					subtt.Fatal(err)
				}
			}()

			_, err = overrideConfigFile.Write(test.json)
			if err != nil {
				subtt.Fatal(err)
			}

			got, err := LoadMediaDbConfig()
			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if !reflect.DeepEqual(test.want, got) {
				subtt.Fatalf("want %v, got %v", test.want, got)
			}
		})
	}
}
