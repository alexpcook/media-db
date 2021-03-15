package schema

type Music struct {
	Title        string `json:"title"`
	Artist       string `json:"artist"`
	YearMade     int    `json:"year"`
	DateListened int64  `json:"date"`
}

func NewMusic(title, artist string, yearMade int, dateListened string) (*Music, error) {
	return &Music{}, nil
}
