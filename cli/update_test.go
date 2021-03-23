package cli

import "testing"

func TestNewUpdateCommand(tt *testing.T) {
	testCases := []struct {
		name    string
		args    []string
		isError bool
	}{
		{"valid-1", []string{"update", "music", "-id", "123", "-title", "title", "-artist", "artist", "-year", "2020", "-date", "2021-01-01"}, false},
		{"valid-2", []string{"update", "movie", "-id", "123", "-title", "title", "-director", "dir", "-year", "2020", "-date", "2021-01-01"}, false},
		{"less-than-two-args", []string{"update"}, true},
		{"invalid-media-type", []string{"update", "invalid"}, true},
		{"invalid-flags-1", []string{"update", "movie", "-notaflag", "movie"}, true},
		{"invalid-flags-2", []string{"update", "music", "-notaflag", "music"}, true},
		{"missing-required-flags-1", []string{"update", "movie", "-title", "movie"}, true},
		{"missing-required-flags-2", []string{"update", "music", "-artist", "artist"}, true},
		{"invalid-value-1", []string{"update", "movie", "-id", "123", "-title", "title", "-director", "director", "-year", "2018", "-date", "bad-date"}, true},
		{"invalid-value-2", []string{"update", "music", "-id", "123", "-title", "title", "-artist", "artist", "-year", "0", "-date", "2021-01-01"}, true},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			_, err := NewUpdateCommand(test.args)

			if test.isError {
				if err == nil {
					subtt.Fatal("want error, got nil")
				}
				return
			} else if err != nil {
				subtt.Fatal(err)
			}
		})
	}
}
