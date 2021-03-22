package cli

import (
	"errors"
	"flag"

	"github.com/alexpcook/media-db-console/schema"
)

type ReadCommand struct {
	FlagSet   *flag.FlagSet
	ID        string
	MediaType schema.Media
}

func NewReadCommand(args []string) (*ReadCommand, error) {
	readCmd := &ReadCommand{}

	// Get everything in the database, so no args or flags are required.
	if len(args) < 2 {
		return readCmd, nil
	}

	mediaType := args[1]

	switch mediaType {
	case MovieMediaType():
		readCmd.FlagSet = flag.NewFlagSet("read movie", flag.ExitOnError)
		readCmd.MediaType = schema.Movie{}
	case MusicMediaType():
		readCmd.FlagSet = flag.NewFlagSet("read music", flag.ExitOnError)
		readCmd.MediaType = schema.Music{}
	default:
		return nil, errors.New(GetInvalidMediaTypeHelpText(ReadCmdName(), mediaType))
	}

	readCmd.FlagSet.StringVar(&readCmd.ID, "id", "", "The id in the database to return")

	err := readCmd.FlagSet.Parse(args[2:])
	if err != nil {
		return nil, err
	}

	return readCmd, nil
}

func (r *ReadCommand) Run() error {
	res, err := MediaDbClient.Read(r.ID, r.MediaType)
	if err != nil {
		return err
	}

	for _, media := range res {
		StdoutLogger.Println(media)
	}

	return nil
}
