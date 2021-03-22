package schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Movie contains information about a single film.
// DateWatched is a Unix timestamp.
type Movie struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Director    string `json:"director"`
	YearMade    int    `json:"year"`
	DateWatched int64  `json:"date"`
}

// Key returns the unique object key for storage in the database.
// For example, /media/movie/6ba7b810-9dad-11d1-80b4-00c04fd430c8
func (m Movie) Key() string {
	return strings.Join([]string{GetBaseKeyFromMediaType(m), m.ID}, "/")
}

// String provides a standard interface to print Movie to output.
func (m Movie) String() string {
	return fmt.Sprintf(`id: %s
  title:    %s
  director: %s
  year:     %d
  date:	    %s`, m.ID, m.Title, m.Director, m.YearMade, time.Unix(m.DateWatched, 0).Format("2006-01-02"))
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
		ID:          uuid.NewString(),
		Title:       title,
		Director:    director,
		YearMade:    yearMade,
		DateWatched: unixTime,
	}, nil
}
