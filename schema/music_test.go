package schema

import "testing"

type newMusicInput struct {
	title  string
	artist string
	year   int
	date   string
}

type newMusicOutput struct {
	music   *Music
	isError bool
}

func TestNewMusic(tt *testing.T) {
}
