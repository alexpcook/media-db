package service

import (
	"context"
	"strings"

	"github.com/alexpcook/media-db/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Delete removes the given media from the database. It returns a non-nil
// error if the entry cannot be deleted.
func (cl *MediaDbClient) Delete(id string, media schema.Media) error {
	objKey := strings.Join([]string{schema.GetBaseKeyFromMediaType(media), id}, "/")

	_, err := cl.s3Client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: &cl.s3Bucket,
		Key:    &objKey,
	})
	if err != nil {
		return err
	}

	return nil
}
