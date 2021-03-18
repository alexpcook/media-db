package schema

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// StringToUnixTime converts a string s of format 'yyyy-mm-dd' to a Unix
// timestamp. The time of the returned value is always 00:00:00 UTC. If
// the string cannot be parsed to a valid date, a non-nil error is returned.
// If s is the empty string "", a Unix time of zero (Jan 1 1970 UTC) is
// returned.
func StringToUnixTime(s string) (int64, error) {
	if s == "" {
		return time.Unix(0, 0).Unix(), nil
	}

	strFields := strings.SplitN(s, "-", 3)
	if numFields := len(strFields); numFields != 3 {
		return 0, fmt.Errorf("length of date string is not 3, got %d", numFields)
	}

	intFields := make([]int, 3)
	var err error
	for i := range intFields {
		intFields[i], err = strconv.Atoi(strFields[i])
		if err != nil {
			return 0, err
		}
	}

	year, month, day := intFields[0], intFields[1], intFields[2]

	if year < 1850 {
		return 0, fmt.Errorf("year must be greater than 1850, got %d", year)
	}

	if month < 1 || month > 12 {
		return 0, fmt.Errorf("month must be between 1 and 12, got %d", month)
	}

	if day < 1 || day > 31 {
		return 0, fmt.Errorf("day must be between 0 and 31, got %d", day)
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC).Unix(), nil
}
