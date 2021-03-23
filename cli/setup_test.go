package cli

import "testing"

func TestNewSetupCommand(tt *testing.T) {
	testCases := []struct {
		name    string
		args    []string
		isError bool
	}{
		{"valid", []string{"setup", "-profile", "prof", "-region", "us-west-1", "-bucket", "my_bucket"}, false},
		{"invalid-flag", []string{"setup", "-notaflag", "test", "-profile", "aws"}, true},
		{"missing-required-flags", []string{"setup", "-profile", "prof"}, true},
		{"invalid-value", []string{"setup", "-profile", "\t", "-region", "us-west-1", "-bucket", "my_bucket"}, true},
	}

	for _, test := range testCases {
		tt.Run(test.name, func(subtt *testing.T) {
			_, err := NewSetupCommand(test.args)

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
