package cli

import "testing"

func TestNewCreateCommand(tt *testing.T) {
	testCases := []struct {
		name    string
		args    []string
		isError bool
	}{
		{"valid-1", []string{"create", "music", "-title", "title", "-artist", "artist", "-year", "2020", "-date", "2021-01-01"}, false},
		{"valid-2", []string{"create", "movie", "-title", "title", "-director", "dir", "-year", "2020", "-date", "2021-01-01"}, false},
		{"less-than-two-args", []string{"create"}, true},
		{"invalid-media-type", []string{"create", "invalid"}, true},
		{"invalid-flags-1", []string{"create", "movie", "-notaflag", "movie"}, true},
		{"invalid-flags-2", []string{"create", "music", "-notaflag", "music"}, true},
		{"missing-required-flags-1", []string{"create", "movie", "-title", "movie"}, true},
		{"missing-required-flags-2", []string{"create", "music", "-artist", "artist"}, true},
		{"invalid-value-1", []string{"create", "movie", "-title", "title", "-director", "director", "-year", "2018", "-date", "bad-date"}, true},
		{"invalid-value-2", []string{"create", "music", "-title", "title", "-artist", "artist", "-year", "0", "-date", "2021-01-01"}, true},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			_, err := NewCreateCommand(test.args)

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
