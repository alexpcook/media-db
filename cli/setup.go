package cli

import (
	"errors"
	"flag"

	"github.com/alexpcook/media-db-console/config"
)

type SetupCommand struct {
	FlagSet *flag.FlagSet
	Config  *config.MediaDbConfig
}

func NewSetupCommand(args []string) (*SetupCommand, error) {
	setupCmd := &SetupCommand{
		FlagSet: flag.NewFlagSet("setup", flag.ExitOnError),
	}

	awsConfig := &config.MediaDbConfig{}
	setupCmd.FlagSet.StringVar(&awsConfig.AWSProfile, "profile", "", "The AWS profile to use")
	setupCmd.FlagSet.StringVar(&awsConfig.AWSRegion, "region", "", "The AWS region to use")
	setupCmd.FlagSet.StringVar(&awsConfig.S3Bucket, "bucket", "", "The S3 bucket to use")

	err := setupCmd.FlagSet.Parse(args[1:])
	if err != nil {
		return nil, err
	}

	expectFlags := 3
	if gotFlags := setupCmd.FlagSet.NFlag(); gotFlags != expectFlags {
		setupCmd.FlagSet.Usage()
		return nil, errors.New("")
	}

	setupCmd.Config, err = config.NewMediaDbConfig(awsConfig.AWSProfile, awsConfig.AWSRegion, awsConfig.S3Bucket)
	if err != nil {
		return nil, err
	}

	return setupCmd, nil
}

func (s *SetupCommand) Run() error {
	return s.Config.Save()
}
