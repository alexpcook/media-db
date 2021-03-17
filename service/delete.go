package service

import (
	"context"

	"github.com/alexpcook/media-db-console/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (cl *MediaDbClient) Delete(media schema.Media) error {
	key := media.S3Key()

	_, err := cl.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Key: &key,
	})
	if err != nil {
		return err
	}

	return nil
}
