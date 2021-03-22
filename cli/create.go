package cli

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/alexpcook/media-db-console/schema"
)

type CreateCommand struct {
	FlagSet  *flag.FlagSet
	NewMedia schema.Media
}

func NewCreateCommand(args []string) (*CreateCommand, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("usage: mdb create %s <flag>... ", strings.Join(GetMediaTypes(), "|"))
	}

	mediaType := args[1]
	createCmd := &CreateCommand{}

	switch mediaType {
	case MovieMediaType():
		createCmd.FlagSet = flag.NewFlagSet("create movie", flag.ExitOnError)
		movie := new(schema.Movie)
		var dateStr string

		createCmd.FlagSet.StringVar(&movie.Title, "title", "", "The title of the movie")
		createCmd.FlagSet.StringVar(&movie.Director, "director", "", "The director of the movie")
		createCmd.FlagSet.IntVar(&movie.YearMade, "year", 0, "The year the movie was made")
		createCmd.FlagSet.StringVar(&dateStr, "date", "", "The date the movie was watched")

		err := createCmd.FlagSet.Parse(args[2:])
		if err != nil {
			return nil, err
		}

		expectFlags := 4
		if gotFlags := createCmd.FlagSet.NFlag(); gotFlags != expectFlags {
			createCmd.FlagSet.Usage()
			return nil, errors.New("")
		}

		createCmd.NewMedia, err = schema.NewMovie(movie.Title, movie.Director, movie.YearMade, dateStr)
		if err != nil {
			return nil, err
		}
	case MusicMediaType():
		createCmd.FlagSet = flag.NewFlagSet("create music", flag.ExitOnError)
		music := new(schema.Music)
		var dateStr string

		createCmd.FlagSet.StringVar(&music.Title, "title", "", "The title of the piece of music")
		createCmd.FlagSet.StringVar(&music.Artist, "artist", "", "The artist who made or performed the piece of music")
		createCmd.FlagSet.IntVar(&music.YearMade, "year", 0, "The year the music was made")
		createCmd.FlagSet.StringVar(&dateStr, "date", "", "The date the music was listened to")

		err := createCmd.FlagSet.Parse(args[2:])
		if err != nil {
			return nil, err
		}

		expectFlags := 4
		if gotFlags := createCmd.FlagSet.NFlag(); gotFlags != expectFlags {
			createCmd.FlagSet.Usage()
			return nil, errors.New("")
		}

		createCmd.NewMedia, err = schema.NewMusic(music.Title, music.Artist, music.YearMade, dateStr)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New(GetInvalidMediaTypeHelpText(CreateCmdName(), mediaType))
	}

	return createCmd, nil
}

func (c *CreateCommand) Run() error {
	return MediaDbClient.Create(c.NewMedia)
}
