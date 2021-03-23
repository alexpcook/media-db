package cli

import (
	"errors"
)

// MediaDbCommand defines methods common to all CLI commands.
type MediaDbCommand interface {
	// Run executes the CLI command and returns any errors encountered during execution.
	Run() error
}

// NewMediaDbCommand parses args from the command line and returns the appropriate
// command for execution. The error will be non-nil if the args do not constitute
// a valid command.
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
