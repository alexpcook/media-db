package schema

// Media is an interface that encompassses methods
// common to all media types in the database.
type Media interface {
	// The unique S3 object key for storage in the database.
	S3Key() string
}
