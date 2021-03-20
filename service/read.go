package service

import (
	"context"
	"encoding/json"
	"io"
	"log"

	"github.com/alexpcook/media-db-console/schema"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Read retrieves the media entries from the database that match the
// specified filters.
func (cl *MediaDbClient) Read(filter string) ([]schema.Media, error) {
	listObjsRes, err := cl.s3Client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &cl.s3Bucket,
		Prefix: &filter,
	})
	if err != nil {
		return nil, err
	}

	mediaRes := make([]schema.Media, 0)

	for _, obj := range listObjsRes.Contents {
		getObjsRes, err := cl.s3Client.GetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: &cl.s3Bucket,
			Key:    obj.Key,
		})
		if err != nil {
			return nil, err
		}

		b := make([]byte, 8)
		jsonData := make([]byte, 0)
		for {
			n, err := getObjsRes.Body.Read(b)
			jsonData = append(jsonData, b[:n]...)
			if err == io.EOF {
				break
			} else if err != nil {
				return nil, err
			}
		}

		media, err := schema.GetMediaTypeFromKey(*obj.Key)
		if err != nil {
			log.Println(err)
		}

		switch media.(type) {
		case schema.Movie:
			movie := schema.Movie{}
			err = json.Unmarshal(jsonData, &movie)
			if err != nil {
				return nil, err
			}
			mediaRes = append(mediaRes, movie)
		case schema.Music:
			music := schema.Music{}
			err = json.Unmarshal(jsonData, &music)
			if err != nil {
				return nil, err
			}
			mediaRes = append(mediaRes, music)
		}
	}

	return mediaRes, nil
}
