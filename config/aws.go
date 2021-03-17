package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

const (
	// DefaultConfigDir is the default AWS configuration directory name.
	DefaultConfigDir string = ".mediadb"
	// DefaultConfigFile is the default AWS configuration file name.
	DefaultConfigFile string = "config"
)

// MediaDbConfig contains the AWS profile and S3 bucket name to use
// for interacting with the database.
type MediaDbConfig struct {
	AWSProfile string `json:"awsprofile"`
	S3Bucket   string `json:"s3bucket"`
}

// GetConfigFilePath returns an absolute filepath to configFileName in
// configDirName in the user's home directory. The error will be non-nil
// if there is a problem determining the user's home directory.
func GetConfigFilePath(configDirName, configFileName string) (string, error) {
	userHomeDirName, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return path.Join(userHomeDirName, configDirName, configFileName), nil
}

// LoadConfig loads the AWS config from filepath and returns a pointer to it.
// The error will be non-nil if the config file cannot be read or its JSON
// cannot be parsed.
func LoadMediaDbConfig(filepath string) (*MediaDbConfig, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var config MediaDbConfig
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(config.AWSProfile) == "" {
		return nil, fmt.Errorf("awsprofile cannot be only space, got %q", config.AWSProfile)
	}

	if strings.TrimSpace(config.S3Bucket) == "" {
		return nil, fmt.Errorf("s3bucket cannot be only space, got %q", config.S3Bucket)
	}

	return &config, nil
}
