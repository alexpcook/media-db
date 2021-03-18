package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

func getDefaultConfigDirName() string {
	return ".mediadb"
}

func getDefaultConfigFileName() string {
	return "config"
}

func getOverrideConfigFilePathEnvVarName() string {
	return "MEDIA_DB_CONFIG_FILE"
}

// MediaDbConfig contains the AWS profile, region, and S3 bucket name to use
// for interacting with the database.
type MediaDbConfig struct {
	AWSProfile string `json:"profile"`
	AWSRegion  string `json:"region"`
	S3Bucket   string `json:"bucket"`
}

func getConfigFilePath() (string, error) {
	if override, isSet := os.LookupEnv(getOverrideConfigFilePathEnvVarName()); isSet {
		_, err := os.Stat(override)
		if err == nil {
			return path.Join(override), nil
		}
	}

	userHomeDirName, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	filepath := path.Join(userHomeDirName, getDefaultConfigDirName(), getDefaultConfigFileName())
	_, err = os.Stat(filepath)
	if err != nil {
		return "", err
	}

	return filepath, nil
}

// LoadMediaDbConfig loads the database config returns a pointer to it. It will
// first look for configuration in a valid override filepath in the
// environment variable MEDIA_DB_CONFIG_FILE. If no valid override file is
// found, it will use the file ./.mediadb/config in the user's home directory.
// The error will be non-nil if a valid config file cannot be found or its
// settings cannot be parsed.
func LoadMediaDbConfig() (*MediaDbConfig, error) {
	configFile, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}

	configData, err := os.ReadFile(configFile)
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
