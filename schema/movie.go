package schema

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// Movie contains information about a single film.
// DateWatched is a Unix timestamp.
type Movie struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	YearMade    int    `json:"year"`
	DateWatched int64  `json:"date"`
}

// S3Key returns the unique S3 object key for storage in the database.
// For example, /movie/2021/QSBNb3ZpZQ==
func (m Movie) S3Key() string {
	baseDir := "movie"

	return strings.Join([]string{
		baseDir,
		fmt.Sprint(m.YearMade),
		base64.URLEncoding.EncodeToString([]byte(m.Title)),
	}, "/")
}

// NewMovie validates the given inputs and returns a pointer to a Movie type.
// The dateWatched parameter should be in the format 'yyyy-mm-dd'.
// If there are validation problems, a non-nil error is returned.
func NewMovie(title, director string, yearMade int, dateWatched string) (*Movie, error) {
	trim := strings.TrimSpace

	title = trim(title)
	if title == "" {
		return nil, fmt.Errorf("title cannot be null, got %q", title)
	}

	director = trim(director)
	if director == "" {
		return nil, fmt.Errorf("director cannot be null, got %q", director)
	}

	if yearMade < 1 {
		return nil, fmt.Errorf("yearMade must be positive, got %d", yearMade)
	}

	unixTime, err := StringToUnixTime(trim(dateWatched))
	if err != nil {
		return nil, err
	}

	return &Movie{
		Title:       title,
		Director:    director,
		YearMade:    yearMade,
		DateWatched: unixTime,
	}, nil
}
