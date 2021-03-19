package service

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/alexpcook/media-db-console/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Add creates or updates a single S3 object in the database.
func (cl *MediaDbClient) Add(media schema.Media) error {
	jsonData, err := json.Marshal(media)
	if err != nil {
		return err
	}

	key := media.Key()
	_, err = cl.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &cl.s3Bucket,
		Key:    &key,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return err
	}

	return nil
}
