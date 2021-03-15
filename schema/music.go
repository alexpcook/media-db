package schema

import "time"

type Music struct {
	Title        string     `json:"title"`
	Artist       string     `json:"artist"`
	YearMade     int        `json:"year"`
	DateListened *time.Time `json:"date"`
}
