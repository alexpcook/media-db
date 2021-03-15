package schema

import (
	"fmt"
	"testing"
	"time"
)

func TestStringToUnixTime(tt *testing.T) {
	testCases := []struct {
		input   string
		isError bool
	}{
		{"2006-01-02T15:04:05Z07:00", false},
		{"2006-01-02", false},
		{"0036-01-02T15:04:05Z07:00", false},
		{"2021-10-30T00:15:15.65Z-08:00", false},
		{"invalid date string", true},
		{"2021-13-11", true},
	}

	unixTimeZero := time.Unix(0, 0)
	for i, test := range testCases {
		tt.Run(fmt.Sprintf("case-%d", i), func(subtt *testing.T) {
			got, err := StringToUnixTime(test.input)

			if test.isError && err == nil {
				subtt.Fatal("want error, got nil")
			}

			if diff := time.Unix(got, 0).Sub(unixTimeZero); diff != time.Duration(got) {
				subtt.Fatalf("want zero diff, got %v", diff)
			}
		})
	}
}
