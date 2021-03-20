package service

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/alexpcook/media-db-console/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Create makes a single new media object in the database.
// It returns a non-nil error if the object cannot be added.
func (cl *MediaDbClient) Create(media schema.Media) error {
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
