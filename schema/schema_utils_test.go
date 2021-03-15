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
		{"2006-01-02T15:04:05Z", false},
		{"2006-01-02T00:00:00-07:00", false},
		{"1900-01-02T15:04:05+07:00", false},
		{"2021-10-30T00:15:15.65Z", false},
		{"invalid date string", true},
		{"2021-11-11", true},
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
