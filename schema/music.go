package schema

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// Music contains information about a single piece of music.
// DateListened is a Unix timestamp.
type Music struct {
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	YearMade     int    `json:"year"`
	DateListened int64  `json:"date"`
}

// S3Key returns the unique S3 object key for storage in the database.
// For example, /music/2021/QW4gQWxidW0=
func (m Music) S3Key() string {
	baseDir := "music"

	return strings.Join([]string{
		baseDir,
		fmt.Sprint(m.YearMade),
		base64.URLEncoding.EncodeToString([]byte(m.Title)),
	}, "/")
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
		Title:        title,
		Artist:       artist,
		YearMade:     yearMade,
		DateListened: unixTime,
	}, nil
}
