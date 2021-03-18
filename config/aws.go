package config

import (
	"encoding/json"
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
