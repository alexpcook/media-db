package schema

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

// Movie contains information about a single film. DateWatched is a Unix timestamp.
type Movie struct {
	Title       string `json:"title"`
	Director    string `json:"director"`
	YearMade    int    `json:"year"`
	DateWatched int64  `json:"date"`
}

// S3Key returns the unique S3 object key for storage in the database.
// For example, /movie/2021/QSBNb3ZpZQ==
func (m *Movie) S3Key() string {
	baseDir := "movie"

	return strings.Join([]string{
		baseDir,
		fmt.Sprint(m.YearMade),
		base64.URLEncoding.EncodeToString([]byte(m.Title)),
	}, "/")
}

// MarshalJSON returns Movie as JSON bytes.
func (m *Movie) MarshalJSON() ([]byte, error) {
	return json.Marshal(*m)
}

// UnmarshalJSON populates the underlying Movie pointer
// with the given JSON bytes b.
func (m *Movie) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, m)
}

// NewMovie validates the fields of a Movie struct, creates it, and returns a pointer.
// The dateWatched parameter should be in the format 'yyyy-mm-dd'.
// If there are validation problems, a non-nil error is returned.
func NewMovie(title, director string, yearMade int, dateWatched string) (*Movie, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, fmt.Errorf("title cannot be only space, got %q", title)
	}

	director = strings.TrimSpace(director)
	if director == "" {
		return nil, fmt.Errorf("director cannot be only space, got %q", director)
	}

	if yearMade < 1 {
		return nil, fmt.Errorf("yearMade must be positive, got %d", yearMade)
	}

	unixTime, err := StringToUnixTime(strings.TrimSpace(dateWatched))
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
