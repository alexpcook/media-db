package cli

import (
	"fmt"
	"strings"
)

func SetupCmdName() string {
	return "setup"
}

func CreateCmdName() string {
	return "create"
}

func ReadCmdName() string {
	return "read"
}

func UpdateCmdName() string {
	return "update"
}

func DeleteCmdName() string {
	return "delete"
}

func MovieMediaType() string {
	return "movie"
}

func MusicMediaType() string {
	return "music"
}

func GetMediaTypes() []string {
	return []string{MovieMediaType(), MusicMediaType()}
}

func GetCLIHelpText() string {
	return fmt.Sprintf(`usage: mdb <command> [%s] [<flag>...]

where <command> is one of:
  setup		Configure the database connection to AWS
  create	Create an entry in the database
  read		Read entries from the database
  update	Update an entry in the database
  delete	Delete an entry from the database`, strings.Join(GetMediaTypes(), "|"))
}

func GetInvalidCommandHelpText(invalidCmd string) string {
	return fmt.Sprintf(`mdb: '%s' is an invalid command

%s`, invalidCmd, GetCLIHelpText())
}

func GetInvalidMediaTypeHelpText(cmd, mediaType string) string {
	return fmt.Sprintf(`mdb: '%s' is an invalid media type, want one of '%s'

%s`, mediaType, strings.Join(GetMediaTypes(), ", "), GetCommandHelpText(cmd))
}

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
