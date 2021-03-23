package cli

import (
	"fmt"
	"strings"
)

// SetupCmdName returns the name of the setup command.
func SetupCmdName() string {
	return "setup"
}

// CreateCmdName returns the name of the create command.
func CreateCmdName() string {
	return "create"
}

// ReadCmdName returns the name of the read command.
func ReadCmdName() string {
	return "read"
}

// UpdateCmdName returns the name of the update command.
func UpdateCmdName() string {
	return "update"
}

// DeleteCmdName returns the name of the delete command.
func DeleteCmdName() string {
	return "delete"
}

// MovieMediaType returns the name of the movie media type.
func MovieMediaType() string {
	return "movie"
}

// MusicMediaType returns the name of the music media type.
func MusicMediaType() string {
	return "music"
}

// GetMediaTypes returns a slice of all valid media types
// that can be stored in the database.
func GetMediaTypes() []string {
	return []string{MovieMediaType(), MusicMediaType()}
}

// GetCLIHelpText returns the general help text for the CLI.
func GetCLIHelpText() string {
	return fmt.Sprintf(`usage: mdb <command> [%s] [<flag>...]

where <command> is one of:
  setup		Configure the database connection to AWS
  create	Create an entry in the database
  read		Read entries from the database
  update	Update an entry in the database
  delete	Delete an entry from the database`, strings.Join(GetMediaTypes(), "|"))
}

// GetInvalidCommandHelpText returns help text intended to be displayed
// when an invalid command is entered by the user.
func GetInvalidCommandHelpText(invalidCmd string) string {
	return fmt.Sprintf(`mdb: '%s' is an invalid command

%s`, invalidCmd, GetCLIHelpText())
}

// GetInvalidMediaTypeHelpText returns help text intended to be displayed
// when an invalid media type is entered by the user.
func GetInvalidMediaTypeHelpText(cmd, mediaType string) string {
	return fmt.Sprintf(`mdb: '%s' is an invalid media type, want one of '%s'

%s`, mediaType, strings.Join(GetMediaTypes(), ", "), GetCommandHelpText(cmd))
}

// GetCommandHelpText returns the usage string for a given command.
func GetCommandHelpText(cmd string) string {
	mediaTypes := strings.Join(GetMediaTypes(), "|")

	switch cmd {
	case CreateCmdName(), UpdateCmdName():
		flagsHelpText := "<flag>..."
		if cmd == UpdateCmdName() {
			flagsHelpText = fmt.Sprintf("%s %s", "-id=<id>", flagsHelpText)
		}
		return fmt.Sprintf(`usage: mdb %s %s %s`, cmd, mediaTypes, flagsHelpText)
	case ReadCmdName():
		return fmt.Sprintf(`usage: mdb %s [%s] [-id=<id>]`, cmd, mediaTypes)
	case DeleteCmdName():
		return fmt.Sprintf(`usage: mdb %s %s -id=<id>`, cmd, mediaTypes)
	default:
		return GetInvalidCommandHelpText(cmd)
	}
}
