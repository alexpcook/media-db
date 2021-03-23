package cli

import "testing"

func TestNewDeleteCommand(tt *testing.T) {
	testCases := []struct {
		name    string
		args    []string
		isError bool
	}{
		{"valid-1", []string{"delete", "movie", "-id", "123"}, false},
		{"valid-2", []string{"delete", "music", "-id", "123"}, false},
		{"less-than-two-args", []string{"delete"}, true},
		{"invalid-media-type", []string{"delete", "invalid"}, true},
		{"invalid-flags-1", []string{"delete", "movie", "-notaflag", "movie"}, true},
		{"invalid-flags-2", []string{"delete", "music", "-notaflag", "music"}, true},
		{"missing-required-flag-1", []string{"delete", "movie"}, true},
		{"missing-required-flag-2", []string{"delete", "music"}, true},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			_, err := NewDeleteCommand(test.args)

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
