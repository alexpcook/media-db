package service

import (
	"context"

	cfg "github.com/alexpcook/media-db-console/config"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// MediaDbClient contains the AWS S3 client and bucket name
// to use for the database.
type MediaDbClient struct {
	s3Client *s3.Client
	s3Bucket string
}

// NewMediaDbClient creates a MediaDbClient from the settings in the given
// MediaDbConfig. Any problem loading the AWS credentials, configuration,
// or accessing the S3 bucket will return a non-nil error.
func NewMediaDbClient(mediaDbConfig *cfg.MediaDbConfig) (*MediaDbClient, error) {
	awsConfig, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(mediaDbConfig.AWSProfile),
		config.WithRegion(mediaDbConfig.AWSRegion))
	if err != nil {
		return nil, err
	}

	s3Client := s3.NewFromConfig(awsConfig)
	mediaDbClient := MediaDbClient{
		s3Client: s3Client,
		s3Bucket: mediaDbConfig.S3Bucket,
	}

	// Validate access to the S3 bucket.
	_, err = mediaDbClient.s3Client.HeadBucket(context.TODO(), &s3.HeadBucketInput{
		Bucket: &mediaDbClient.s3Bucket,
	})
	if err != nil {
		return nil, err
	}

	return &mediaDbClient, nil
}
