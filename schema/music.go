package schema

import "time"

type Music struct {
	Title        string    `json:"title"`
	Artist       string    `json:"artist"`
	YearMade     int       `json:"year"`
	DateListened time.Time `json:"date"`
}

func NewMusic(title, artist string, yearMade int, dateListened string) (*Music, error) {
	return &Music{}, nil
}
