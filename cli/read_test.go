package cli

import "testing"

func TestNewReadCommand(tt *testing.T) {
	testCases := []struct {
		name    string
		args    []string
		isError bool
	}{
		{"valid-1", []string{"read"}, false},
		{"valid-2", []string{"read", "movie"}, false},
		{"valid-3", []string{"read", "music"}, false},
		{"valid-4", []string{"read", "movie", "-id", "123"}, false},
		{"valid-5", []string{"read", "music", "-id", "123"}, false},
		{"invalid-media-type", []string{"read", "invalid"}, true},
		{"invalid-flags-1", []string{"read", "movie", "-notaflag", "movie"}, true},
		{"invalid-flags-2", []string{"read", "music", "-notaflag", "music"}, true},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			_, err := NewReadCommand(test.args)

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
