package schema

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMovieJsonUnmarshal(t *testing.T) {
	title := "A Movie Title"
	director := "A Movie Director"
	year := 2021
	date := "2021-03-14T13:14:15.32-08:00" // string must conform to RFC 3339

	movieBytes := []byte(fmt.Sprintf(`{
		"title": "%s",
		"director": "%s",
		"year": %d,
		"date": "%s"
	}`, title, director, year, date))
	movieStruct := Movie{}

	err := json.Unmarshal(movieBytes, &movieStruct)
	if err != nil {
		t.Fatal(err)
	}

	if movieStruct.Title != title {
		t.Fatalf("Expected %s for Movie.Title field, got %s", title, movieStruct.Title)
	}

	if movieStruct.Director != director {
		t.Fatalf("Expected %s for Movie.Director field, got %s", director, movieStruct.Director)
	}

	if movieStruct.YearMade != year {
		t.Fatalf("Expected %d for Movie.YearMade field, got %d", year, movieStruct.YearMade)
	}

	if d, err := time.Parse(time.RFC3339, date); !d.Equal(*movieStruct.DateWatched) || err != nil {
		t.Fatalf("Expected %s for Movie.DateWatched field, got %s %v", date, movieStruct.DateWatched, err)
	}
}
