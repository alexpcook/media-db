package cli

import (
	"errors"
	"flag"
	"fmt"
	"strings"

	"github.com/alexpcook/media-db-console/schema"
)

type UpdateCommand struct {
	FlagSet      *flag.FlagSet
	ID           string
	UpdatedMedia schema.Media
}

func NewUpdateCommand(args []string) (*UpdateCommand, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("usage: mdb udpate %s <flag>... ", strings.Join(GetMediaTypes(), "|"))
	}

	mediaType := args[1]
	updateCmd := &UpdateCommand{}

	switch mediaType {
	case MovieMediaType():
		updateCmd.FlagSet = flag.NewFlagSet("update movie", flag.ExitOnError)
		movie := new(schema.Movie)
		var dateStr string

		updateCmd.FlagSet.StringVar(&updateCmd.ID, "id", "", "The id of the movie to udpate")
		updateCmd.FlagSet.StringVar(&movie.Title, "title", "", "The title of the movie")
		updateCmd.FlagSet.StringVar(&movie.Director, "director", "", "The director of the movie")
		updateCmd.FlagSet.IntVar(&movie.YearMade, "year", 0, "The year the movie was made")
		updateCmd.FlagSet.StringVar(&dateStr, "date", "", "The date the movie was watched")

		err := updateCmd.FlagSet.Parse(args[2:])
		if err != nil {
			return nil, err
		}

		expectFlags := 4
		if gotFlags := updateCmd.FlagSet.NFlag(); gotFlags != expectFlags {
			updateCmd.FlagSet.Usage()
			return nil, errors.New("")
		}

		movie, err = schema.NewMovie(movie.Title, movie.Director, movie.YearMade, dateStr)
		if err != nil {
			return nil, err
		}

		movie.ID = updateCmd.ID
		updateCmd.UpdatedMedia = *movie
	case MusicMediaType():
		updateCmd.FlagSet = flag.NewFlagSet("update music", flag.ExitOnError)
		music := new(schema.Music)
		var dateStr string

		updateCmd.FlagSet.StringVar(&updateCmd.ID, "id", "", "The id of the music to udpate")
		updateCmd.FlagSet.StringVar(&music.Title, "title", "", "The title of the piece of music")
		updateCmd.FlagSet.StringVar(&music.Artist, "artist", "", "The artist who made or performed the piece of music")
		updateCmd.FlagSet.IntVar(&music.YearMade, "year", 0, "The year the music was made")
		updateCmd.FlagSet.StringVar(&dateStr, "date", "", "The date the music was listened to")

		err := updateCmd.FlagSet.Parse(args[2:])
		if err != nil {
			return nil, err
		}

		expectFlags := 4
		if gotFlags := updateCmd.FlagSet.NFlag(); gotFlags != expectFlags {
			updateCmd.FlagSet.Usage()
			return nil, errors.New("")
		}

		music, err = schema.NewMusic(music.Title, music.Artist, music.YearMade, dateStr)
		if err != nil {
			return nil, err
		}

		music.ID = updateCmd.ID
		updateCmd.UpdatedMedia = *music
	default:
		return nil, errors.New(GetInvalidMediaTypeHelpText(UpdateCmdName(), mediaType))
	}

	return updateCmd, nil
}

func (u *UpdateCommand) Run() error {
	return MediaDbClient.Update(u.ID, u.UpdatedMedia)
}
