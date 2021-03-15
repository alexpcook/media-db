package schema

import "testing"

type newMovieInput struct {
	title    string
	director string
	year     int
	date     string
}

type newMovieOutput struct {
	movie   *Movie
	isError bool
}

func TestNewMovie(tt *testing.T) {
}
