package schema

import (
	"reflect"
	"strings"
	"testing"

	"github.com/google/uuid"
)

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
	testCases := []struct {
		name   string
		input  newMusicInput
		output newMusicOutput
	}{
		{
			"basic",
			newMusicInput{"a title", "an artist", 2000, "2021-03-14"},
			newMusicOutput{&Music{}, false},
		},
		{
			"empty title",
			newMusicInput{"", "an artist", 2000, "2021-03-14"},
			newMusicOutput{nil, true},
		},
		{
			"empty artist",
			newMusicInput{"a title", "\t  \t\n", 2000, "2021-03-14"},
			newMusicOutput{nil, true},
		},
		{
			"invalid year",
			newMusicInput{"a title", "an artist", -100, "2021-03-14"},
			newMusicOutput{nil, true},
		},
		{
			"invalid date",
			newMusicInput{"a title", "an artist", 2000, "2021-17-14"},
			newMusicOutput{nil, true},
		},
	}

	testUUID := uuid.NewString()

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			music, err := NewMusic(test.input.title, test.input.artist, test.input.year, test.input.date)
			if music != nil {
				music.ID = testUUID // force the UUID to be constant for testing purposes
			}

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

			test.output.music.ID = testUUID
			test.output.music.Title = test.input.title
			test.output.music.Artist = test.input.artist
			test.output.music.YearMade = test.input.year
			test.output.music.DateListened = t

			if !reflect.DeepEqual(test.output.music, music) {
				subtt.Fatalf("want %v, got %v", test.output.music, music)
			}

			wantKey := strings.Join([]string{GetMediaBaseKey(), GetMediaTypeKey(*music), testUUID}, "/")
			if gotKey := music.Key(); wantKey != gotKey {
				subtt.Fatalf("s3 key error: want %v, got %v", wantKey, gotKey)
			}
		})
	}
}
