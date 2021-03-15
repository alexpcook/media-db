package schema

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMusicJsonMarshal(t *testing.T) {
	testTime, err := time.Parse(time.RFC3339, "2008-10-29T01:56:19.32-01:00")
	if err != nil {
		t.Fatal(err)
	}

	musicStruct := Music{
		Title:        "An Album Title",
		Artist:       "An Artist",
		YearMade:     1993,
		DateListened: &testTime,
	}

	_, err = json.Marshal(musicStruct)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMusicJsonUnmarshal(t *testing.T) {
	title := "An Album Title"
	artist := "An Artist"
	year := 1983
	date := "2008-10-29T01:56:19.32-01:00" // string must conform to RFC 3339

	musicBytes := []byte(fmt.Sprintf(`{
		"title": "%s",
		"artist": "%s",
		"year": %d,
		"date": "%s"
	}`, title, artist, year, date))
	musicStruct := Music{}

	err := json.Unmarshal(musicBytes, &musicStruct)
	if err != nil {
		t.Fatal(err)
	}

	if musicStruct.Title != title {
		t.Fatalf("Expected %s for Music.Title field, got %s", title, musicStruct.Title)
	}

	if musicStruct.Artist != artist {
		t.Fatalf("Expected %s for Music.Artist field, got %s", artist, musicStruct.Artist)
	}

	if musicStruct.YearMade != year {
		t.Fatalf("Expected %d for Music.YearMade field, got %d", year, musicStruct.YearMade)
	}

	if d, err := time.Parse(time.RFC3339, date); !d.Equal(*musicStruct.DateListened) || err != nil {
		t.Fatalf("Expected %s for Music.DateListened field, got %s %v", date, musicStruct.DateListened, err)
	}
}
