package schema

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getMediaKey() string {
	return "media"
}

func getMusicKey() string {
	return "music"
}

func getMovieKey() string {
	return "movie"
}

func getUnknownKey() string {
	return "unknown"
}

// GetBaseKeyFromMediaType returns the base key string associated
// with a particular type of media. A concrete media type appends
// a UUID onto this base key with its Key() method before storage
// in the database.
func GetBaseKeyFromMediaType(media Media) string {
	var typeKey string

	switch media.(type) {
	case Movie:
		typeKey = getMovieKey()
	case Music:
		typeKey = getMusicKey()
	default:
		typeKey = getUnknownKey()
	}

	return strings.Join([]string{getMediaKey(), typeKey}, "/")
}

// GetMediaTypeFromKey returns a concrete type that implements the media interface
// given a key string from the database. It will return a non-nil error if the key
// string cannot be determined to be a valid type of media.
func GetMediaTypeFromKey(key string) (Media, error) {
	keyParts := strings.Split(key, "/")
	baseKey := strings.Join(keyParts[:len(keyParts)-1], "/")

	var movie Movie
	var music Music

	switch baseKey {
	case GetBaseKeyFromMediaType(movie):
		return movie, nil
	case GetBaseKeyFromMediaType(music):
		return music, nil
	default:
		return nil, fmt.Errorf("key %s does not correspond to a valid media type", key)
	}
}

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
