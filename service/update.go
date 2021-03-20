package service

import "github.com/alexpcook/media-db-console/schema"

// Update changes a single existing media object in the database.
// It returns a non-nil error if the object cannot be updated.
func (cl *MediaDbClient) Update(id string, media schema.Media) error {
	return nil
}
