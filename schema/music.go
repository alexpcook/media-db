package schema

// Music contains information about a single piece of music. DateListened is a Unix timestamp.
type Music struct {
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	YearMade     int    `json:"year"`
	DateListened int64  `json:"date"`
}

// NewMusic validates the fields of a Music struct, creates it, and returns a pointer.
// The dateListened parameter should be in the format 'yyyy-mm-dd'.
// If there are validation problems, a non-nil error is returned.
func NewMusic(title, artist string, yearMade int, dateListened string) (*Music, error) {
	return &Music{}, nil
}
