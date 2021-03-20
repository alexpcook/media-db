package service

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"

	"github.com/alexpcook/media-db-console/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Update changes a single existing media object in the database.
// It returns a non-nil error if the object cannot be updated.
func (cl *MediaDbClient) Update(id string, newMedia schema.Media) error {
	objKey := strings.Join([]string{schema.GetBaseKeyFromMediaType(newMedia), id}, "/")

	// Validate that the object exists (don't create it if it doesn't).
	_, err := cl.s3Client.HeadObject(context.TODO(), &s3.HeadObjectInput{
		Bucket: &cl.s3Bucket,
		Key:    &objKey,
	})
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(newMedia)
	if err != nil {
		return err
	}

	_, err = cl.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &cl.s3Bucket,
		Key:    &objKey,
		Body:   bytes.NewReader(jsonData),
	})
	if err != nil {
		return err
	}

	return nil
}
