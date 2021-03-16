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

func TestLoadConfig(tt *testing.T) {
	testCases := []struct {
		name, aws, s3 string
		config        *MediaDbConfig
		isError       bool
	}{
		{"basic", "test-profile", "test-bucket", &MediaDbConfig{}, false},
		{"invalid-filepath", "test-profile", "test-bucket", nil, true},
		{"invalid-json", "test-profile", "test-bucket", nil, true},
		{"null-profile", "   ", "test-bucket", nil, true},
		{"null-bucket", "test-profile", "", nil, true},
	}

	for i, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			data := []byte(fmt.Sprintf(`{"awsprofile": "%s", "s3bucket": "%s"}`, test.aws, test.s3))
			if i == 2 { // invalid json test case
				data = []byte(fmt.Sprintf(`{"awsprofile: "%s", "s3bucket": "%s"}`, test.aws, test.s3))
			}

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

			_, err = tempConfigFile.Write(data)
			if err != nil {
				subtt.Fatal(err)
			}

			filepath := tempConfigFile.Name()
			if i == 1 { // invalid filepath test case
				filepath = "an invalid filepath name"
			}

			got, err := LoadConfig(filepath)
			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			test.config.AWSProfile = test.aws
			test.config.S3BucketName = test.s3

			if !reflect.DeepEqual(test.config, got) {
				subtt.Fatalf("want %v, got %v", test.config, got)
			}
		})
	}
}
