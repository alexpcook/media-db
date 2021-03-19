package schema

import (
	"encoding/base64"
	"fmt"
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
		{
			"empty title",
			newMovieInput{"", "a director", 2000, "2021-03-14"},
			newMovieOutput{nil, true},
		},
		{
			"empty director",
			newMovieInput{"a title", "\t  \t\n", 2000, "2021-03-14"},
			newMovieOutput{nil, true},
		},
		{
			"invalid year",
			newMovieInput{"a title", "a director", 0, "2021-03-14"},
			newMovieOutput{nil, true},
		},
		{
			"invalid date",
			newMovieInput{"a title", "a director", 2000, "2021-17-14"},
			newMovieOutput{nil, true},
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

			wantKey := "movie" + "/" + fmt.Sprint(movie.YearMade) + "/" + base64.URLEncoding.EncodeToString([]byte(movie.Title))
			if gotKey := movie.Key(); wantKey != gotKey {
				subtt.Fatalf("s3 key error: want %v, got %v", wantKey, gotKey)
			}
		})
	}
}
