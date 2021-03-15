package schema

import "time"

type Movie struct {
	Title       string     `json:"title"`
	Director    string     `json:"director"`
	YearMade    int        `json:"year"`
	DateWatched *time.Time `json:"date"`
}

func NewMovie(title, director string, yearMade int, dateWatched string) (*Movie, error) {
	return &Movie{}, nil
}
