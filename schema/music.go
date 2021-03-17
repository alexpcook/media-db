package schema

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// Music contains information about a single piece of music. DateListened is a Unix timestamp.
type Music struct {
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	YearMade     int    `json:"year"`
	DateListened int64  `json:"date"`
}

func (m *Music) S3Key() string {
	baseDir := "music"

	return strings.Join([]string{
		baseDir,
		fmt.Sprint(m.YearMade),
		base64.URLEncoding.EncodeToString([]byte(m.Title)),
	}, "/")
}

// NewMusic validates the fields of a Music struct, creates it, and returns a pointer.
// The dateListened parameter should be in the format 'yyyy-mm-dd'.
// If there are validation problems, a non-nil error is returned.
func NewMusic(title, artist string, yearMade int, dateListened string) (*Music, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return nil, fmt.Errorf("title cannot be only space, got %q", title)
	}

	artist = strings.TrimSpace(artist)
	if artist == "" {
		return nil, fmt.Errorf("artist cannot be only space, got %q", artist)
	}

	if yearMade < 1 {
		return nil, fmt.Errorf("yearMade must be positive, got %d", yearMade)
	}

	unixTime, err := StringToUnixTime(strings.TrimSpace(dateListened))
	if err != nil {
		return nil, err
	}

	return &Music{
		Title:        title,
		Artist:       artist,
		YearMade:     yearMade,
		DateListened: unixTime,
	}, nil
}
