package schema

import "fmt"

// Media is an interface that encompassses methods
// common to all media types in the database.
type Media interface {
	// Key returns the unique object string for storage in the database.
	Key() string

	// Each media type should specify a unique way to be printed.
	fmt.Stringer
}
