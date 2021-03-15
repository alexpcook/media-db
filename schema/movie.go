package schema

import (
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

// NewMovie validates the fields of a Movie struct, creates it, and returns a pointer.
// If there are validation problems, a non-nil error is returned.
func NewMovie(title, director string, yearMade int, dateWatched string) (*Movie, error) {
	if strings.TrimSpace(title) == "" {
		return nil, fmt.Errorf("title cannot be only space, got %q", title)
	}

	if strings.TrimSpace(director) == "" {
		return nil, fmt.Errorf("director cannot be only space, got %q", director)
	}

	if yearMade < 1 {
		return nil, fmt.Errorf("yearMade must be positive, got %d", yearMade)
	}

	unixTime, err := StringToUnixTime(dateWatched)
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
