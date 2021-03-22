package cli

import (
	"errors"
)

type MediaDbCommand interface {
	Run() error
}

func NewMediaDbCommand(args []string) (MediaDbCommand, error) {
	switch cmd := args[0]; cmd {
	case SetupCmdName():
		return NewSetupCommand(args)
	case CreateCmdName():
		InitDb()
		return NewCreateCommand(args)
	case ReadCmdName():
		InitDb()
		return NewReadCommand(args)
	case UpdateCmdName():
		InitDb()
		return NewUpdateCommand(args)
	case DeleteCmdName():
		InitDb()
		return NewDeleteCommand(args)
	default:
		return nil, errors.New(GetInvalidCommandHelpText(cmd))
	}
}
