package schema

import (
	"reflect"
	"testing"
)

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
	testCases := []struct {
		name   string
		input  newMovieInput
		output newMovieOutput
	}{
		{
			"basic",
			newMovieInput{"a title", "a director", 2000, "2021-03-14"},
			newMovieOutput{&Movie{}, false},
		},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			movie, err := NewMovie(test.input.title, test.input.director, test.input.year, test.input.date)

			if test.output.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			t, err := StringToUnixTime(test.input.date)
			if err != nil {
				subtt.Fatal(err)
			}

			test.output.movie.Title = test.input.title
			test.output.movie.Director = test.input.director
			test.output.movie.YearMade = test.input.year
			test.output.movie.DateWatched = t

			if !reflect.DeepEqual(test.output.movie, movie) {
				subtt.Fatalf("want %v, got %v", test.output.movie, movie)
			}
		})
	}
}
