package config

import (
	"errors"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
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

func TestGetDefaultConfigFile(tt *testing.T) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		tt.Fatal(err)
	}

	want := path.Join(homeDir, ".mediadb", "config")

	if got := GetDefaultConfigFile(); want != got {
		tt.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetOverrideConfigFile(tt *testing.T) {
	testCases := []string{path.Join("some", "override", "file"), ""}

	for _, want := range testCases {
		var err error

		if want != "" {
			err = os.Setenv(GetOverrideConfigFileEnvVar(), want)
		} else {
			err = os.Unsetenv(GetOverrideConfigFileEnvVar())
		}

		if err != nil {
			tt.Fatal(err)
		}

		if got := GetOverrideConfigFile(); want != got {
			tt.Fatalf("want %q, got %q", want, got)
		}
	}
}

func TestGetOverrideConfigFileEnvVar(tt *testing.T) {
	want := "MEDIA_DB_CONFIG_FILE"

	if got := GetOverrideConfigFileEnvVar(); want != got {
		tt.Fatalf("want %q, got %q", want, got)
	}
}

func TestGetCurrentConfigFile(tt *testing.T) {
	testCases := []string{path.Join("some", "override", "file"), ""}

	for _, file := range testCases {
		var err error
		var want string

		if file != "" {
			err = os.Setenv(GetOverrideConfigFileEnvVar(), file)
			want = file
		} else {
			err = os.Unsetenv(GetOverrideConfigFileEnvVar())
			want = GetDefaultConfigFile()
		}

		if err != nil {
			tt.Fatal(err)
		}

		if got := GetCurrentConfigFile(); want != got {
			tt.Fatalf("want %q, got %q", want, got)
		}
	}
}

func TestNewMediaDbConfig(tt *testing.T) {
	testCases := []struct {
		name                    string
		profile, region, bucket string
		want                    *MediaDbConfig
		isError                 bool
	}{
		{
			"valid",
			"aws-profile", "aws-region", "aws-bucket",
			&MediaDbConfig{"aws-profile", "aws-region", "aws-bucket"},
			false,
		},
		{
			"null-profile",
			"\t \n", "aws-region", "aws-bucket",
			nil,
			true,
		},
		{
			"null-region",
			"aws-profile", "    ", "aws-bucket",
			nil,
			true,
		},
		{
			"null-bucket",
			"aws-profile", "aws-region", "",
			nil,
			true,
		},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			got, err := NewMediaDbConfig(test.profile, test.region, test.bucket)

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

func TestMediaDbConfigSave(tt *testing.T) {
	defer func() {
		err := os.Unsetenv(GetOverrideConfigFileEnvVar())
		if err != nil {
			tt.Fatal(err)
		}
	}()

	defer func(filename string) {
		err := os.Remove(filename)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			tt.Fatal(err)
		}
	}(GetCurrentConfigFile())

	cfg, err := NewMediaDbConfig("profile", "region", "bucket")
	if err != nil {
		tt.Fatal(err)
	}

	err = cfg.Save()
	if err != nil {
		tt.Fatal(err)
	}

	err = os.Setenv(GetOverrideConfigFileEnvVar(), path.Join(os.TempDir(), ".media_db_test", "config"))
	if err != nil {
		tt.Fatal(err)
	}

	err = cfg.Save()
	if err != nil {
		tt.Fatal(err)
	}

	val, exists := os.LookupEnv(GetOverrideConfigFileEnvVar())
	if !exists {
		tt.Fatalf("expected %s to be set, but was not", GetOverrideConfigFileEnvVar())
	}
	defer func(filename string) {
		if !strings.Contains(filename, os.TempDir()) {
			tt.Fatalf("unexpected value of testing directory, got %q", filename)
		}
		err = os.RemoveAll(path.Dir(filename))
		if err != nil {
			tt.Fatal(err)
		}
	}(val)

	tempFile, err := os.CreateTemp(os.TempDir(), "media_db_test")
	if err != nil {
		tt.Fatal(err)
	}
	defer func() {
		err = os.Remove(tempFile.Name())
		if err != nil {
			tt.Fatal(err)
		}
	}()

	err = os.Setenv(GetOverrideConfigFileEnvVar(), path.Join(tempFile.Name(), "config"))
	if err != nil {
		tt.Fatal(err)
	}

	// This should result in an error since the parent of the config file path is not a directory.
	err = cfg.Save()
	if err == nil {
		tt.Fatal("want error, got nil")
	}
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
