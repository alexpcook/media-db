package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

// GetDefaultConfigFile returns the default database configuration file
// location. For example, $HOME/.mediadb/config on Unix systems.
func GetDefaultConfigFile() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return path.Join(homeDir, ".mediadb", "config")
}

// GetOverrideConfigFile returns the override database configuration file
// location if it is set via the environment variable returned by
// GetOverrideConfigFileEnvVar, else "".
func GetOverrideConfigFile() string {
	if override, isSet := os.LookupEnv(GetOverrideConfigFileEnvVar()); isSet {
		return path.Join(override)
	}

	return ""
}

// GetOverrideConfigFileEnvVar returns the name of the environment variable
// used to override the default database configuration file location.
func GetOverrideConfigFileEnvVar() string {
	return "MEDIA_DB_CONFIG_FILE"
}

// GetCurrentConfigFile returns the current database configuration file.
func GetCurrentConfigFile() string {
	if overrideFile := GetOverrideConfigFile(); overrideFile != "" {
		return overrideFile
	}

	return GetDefaultConfigFile()
}

// MediaDbConfig contains the AWS profile, region, and S3 bucket name to use
// for interacting with the database.
type MediaDbConfig struct {
	AWSProfile string `json:"profile"`
	AWSRegion  string `json:"region"`
	S3Bucket   string `json:"bucket"`
}

// NewMediaDbConfig returns a pointer based on the given AWS profile,
// AWS region, and S3 bucket.
func NewMediaDbConfig(awsProfile, awsRegion, s3Bucket string) (*MediaDbConfig, error) {
	trim := strings.TrimSpace

	awsProfile = trim(awsProfile)
	if awsProfile == "" {
		return nil, fmt.Errorf("awsProfile cannot be null, got %q", awsProfile)
	}

	awsRegion = trim(awsRegion)
	if awsRegion == "" {
		return nil, fmt.Errorf("awsRegion cannot be null, got %q", awsRegion)
	}

	s3Bucket = trim(s3Bucket)
	if s3Bucket == "" {
		return nil, fmt.Errorf("s3Bucket cannot be null, got %q", s3Bucket)
	}

	return &MediaDbConfig{
		AWSProfile: awsProfile,
		AWSRegion:  awsRegion,
		S3Bucket:   s3Bucket,
	}, nil
}

// Save creates the MediaDbConfig as JSON in the
// current configuration file location.
func (cfg *MediaDbConfig) Save() error {
	configFile := GetCurrentConfigFile()

	configFileNew, err := os.CreateTemp("", "media_db_config")
	if err != nil {
		return err
	}
	defer func() {
		err = os.Remove(configFileNew.Name())
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			log.Fatal(err)
		}
	}()

	configData, err := json.Marshal(*cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFileNew.Name(), configData, 0600)
	if err != nil {
		return err
	}

	err = os.Rename(configFileNew.Name(), configFile)
	if err != nil {
		return err
	}

	return nil
}

// LoadMediaDbConfig loads the database config and returns a pointer to it. It
// will first look for configuration in an override file in the environment
// variable MEDIA_DB_CONFIG_FILE. If no override file is found, it will use the
// file ~/.mediadb/config in the user's home directory. The error will be
// non-nil if a valid config file cannot be found or its settings cannot be parsed.
func LoadMediaDbConfig() (*MediaDbConfig, error) {
	configData, err := os.ReadFile(GetCurrentConfigFile())
	if err != nil {
		return nil, err
	}

	dbConfig := MediaDbConfig{}
	err = json.Unmarshal(configData, &dbConfig)
	if err != nil {
		return nil, err
	}

	trim := strings.TrimSpace
	if trim(dbConfig.AWSProfile) == "" {
		return nil, fmt.Errorf("profile cannot be null, got %q", dbConfig.AWSProfile)
	}

	if trim(dbConfig.AWSRegion) == "" {
		return nil, fmt.Errorf("region cannot be null, got %q", dbConfig.AWSRegion)
	}

	if trim(dbConfig.S3Bucket) == "" {
		return nil, fmt.Errorf("bucket cannot be null, got %q", dbConfig.S3Bucket)
	}

	return &dbConfig, nil
}
