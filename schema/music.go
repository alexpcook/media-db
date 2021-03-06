package schema

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Music contains information about a single piece of music.
// DateListened is a Unix timestamp.
type Music struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	YearMade     int    `json:"year"`
	DateListened int64  `json:"date"`
}

// Key returns the unique object key for storage in the database.
// For example, /media/music/6ba7b810-9dad-11d1-80b4-00c04fd430c8
func (m Music) Key() string {
	return strings.Join([]string{GetBaseKeyFromMediaType(m), m.ID}, "/")
}

// String provides a standard interface to print Music to output.
func (m Music) String() string {
	return fmt.Sprintf(`id: %s
  title:  %s
  artist: %s
  year:   %d
  date:	  %s`, m.ID, m.Title, m.Artist, m.YearMade, time.Unix(m.DateListened, 0).Format("2006-01-02"))
}

// NewMusic validates the given inputs and returns a pointer to a Music type.
// The dateListened parameter should be in the format 'yyyy-mm-dd'.
// If there are validation problems, a non-nil error is returned.
func NewMusic(title, artist string, yearMade int, dateListened string) (*Music, error) {
	trim := strings.TrimSpace

	title = trim(title)
	if title == "" {
		return nil, fmt.Errorf("title cannot be null, got %q", title)
	}

	artist = trim(artist)
	if artist == "" {
		return nil, fmt.Errorf("artist cannot be null, got %q", artist)
	}

	if yearMade < 1 {
		return nil, fmt.Errorf("yearMade must be positive, got %d", yearMade)
	}

	unixTime, err := StringToUnixTime(trim(dateListened))
	if err != nil {
		return nil, err
	}

	return &Music{
		ID:           uuid.NewString(),
		Title:        title,
		Artist:       artist,
		YearMade:     yearMade,
		DateListened: unixTime,
	}, nil
}
