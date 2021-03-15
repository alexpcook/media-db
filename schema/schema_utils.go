package schema

import (
	"time"
)

// StringToUnixTime converts a string s that meets the specifications
// of RFC 3339 to a Unix timestamp. If the string cannot be parsed
// according to RFC 3339, a non-nil parsing error is returned.
func StringToUnixTime(s string) (int64, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return 0, err
	}

	return t.Unix(), nil
}
