package cli

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/alexpcook/media-db-console/schema"
)

// DeleteCommand provides an interface between the CLI and the MediaDbClient delete service.
type DeleteCommand struct {
	FlagSet   *flag.FlagSet
	ID        string
	MediaType schema.Media
}

// NewDeleteCommand returns a pointer to a new DeleteCommand struct. If there is a problem
// creating the command, the usage help text for the command will be returned as a non-nil error.
func NewDeleteCommand(args []string) (*DeleteCommand, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("usage: mdb delete %s -id=<id>", strings.Join(GetMediaTypes(), "|"))
	}

	mediaType := args[1]
	deleteCmd := &DeleteCommand{}

	switch mediaType {
	case MovieMediaType():
		deleteCmd.FlagSet = flag.NewFlagSet("delete movie", flag.ExitOnError)
		deleteCmd.MediaType = schema.Movie{}
	case MusicMediaType():
		deleteCmd.FlagSet = flag.NewFlagSet("delete music", flag.ExitOnError)
		deleteCmd.MediaType = schema.Music{}
	default:
		return nil, errors.New(GetInvalidMediaTypeHelpText(DeleteCmdName(), mediaType))
	}

	deleteCmd.FlagSet.StringVar(&deleteCmd.ID, "id", "", "The id in the database to delete")

	err := deleteCmd.FlagSet.Parse(args[2:])
	if err != nil {
		return nil, err
	}

	expectFlags := 1
	if gotFlags := deleteCmd.FlagSet.NFlag(); gotFlags != expectFlags {
		deleteCmd.FlagSet.Usage()
		return nil, errors.New("")
	}

	return deleteCmd, nil
}

// Run executes the DeleteCommand. It returns a non-nil error
// if the underlying delete service encounters a problem.
func (d *DeleteCommand) Run() error {
	return MediaDbClient.Delete(d.ID, d.MediaType)
}
