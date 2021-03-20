package schema

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestGetBaseKeyFromMediaType(tt *testing.T) {
	testCases := []struct {
		media Media
		want  string
	}{
		{Movie{}, getMediaKey() + "/" + getMovieKey()},
		{Music{}, getMediaKey() + "/" + getMusicKey()},
		{nil, getMediaKey() + "/" + getUnknownKey()},
	}

	for _, test := range testCases {
		if got := GetBaseKeyFromMediaType(test.media); test.want != got {
			tt.Fatalf("for type %T, want %s, got %s", test.media, test.want, got)
		}
	}
}

func TestGetMediaTypeFromKey(tt *testing.T) {
	testCases := []struct {
		key     string
		want    Media
		isError bool
	}{
		{getMediaKey() + "/" + getMovieKey() + "/" + uuid.NewString(), Movie{}, false},
		{getMediaKey() + "/" + getMusicKey() + "/" + uuid.NewString(), Music{}, false},
		{getMediaKey() + "/" + getUnknownKey() + "/" + uuid.NewString(), nil, true},
		{uuid.NewString(), nil, true},
	}

	for _, test := range testCases {
		got, err := GetMediaTypeFromKey(test.key)

		if test.isError {
			if err == nil {
				tt.Fatal("want error, got nil")
			}
		} else if err != nil {
			tt.Fatal(err)
		}

		if !reflect.DeepEqual(test.want, got) {
			tt.Fatalf("want %v, got %v", test.want, got)
		}
	}
}

func TestStringToUnixTime(tt *testing.T) {
	testCases := []struct {
		input   string
		isError bool
	}{
		{"", false},
		{"2006-01-02", false},
		{"2021-03-14", false},
		{"1990-06-12", false},
		{"2021", true},
		{"2021-05", true},
		{"abc-05-21", true},
		{"2021--21", true},
		{"2021-04-1.4", true},
		{"0-01-02", true},
		{"1851-01-02", false},
		{"2006-0-04", true},
		{"2006-1-04", false},
		{"2006-13-02", true},
		{"2006-12-02", false},
		{"2006-06-00", true},
		{"2006-06-01", false},
		{"2006-06-32", true},
		{"2006-06-31", false},
	}

	unixTimeZero := time.Unix(0, 0)
	for i, test := range testCases {
		tt.Run(fmt.Sprintf("case-%d", i), func(subtt *testing.T) {
			got, err := StringToUnixTime(test.input)

			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}

			if diff := time.Unix(got, 0).Sub(unixTimeZero) / 1e9; diff != time.Duration(got) {
				subtt.Fatalf("want zero diff, got %v", diff)
			}
		})
	}
}
