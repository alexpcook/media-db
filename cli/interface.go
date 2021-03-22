package cli

import (
	"errors"
)

type MediaDbCommand interface {
	Run() error
}

func NewMediaDbCommand(args []string) (MediaDbCommand, error) {
	switch cmd := args[0]; cmd {
	case CreateCmdName():
		return NewCreateCommand(args)
	case ReadCmdName():
		return NewReadCommand(args)
	case UpdateCmdName():
		return NewUpdateCommand(args)
	case DeleteCmdName():
		return NewDeleteCommand(args)
	default:
		return nil, errors.New(GetInvalidCommandHelpText(cmd))
	}
}
