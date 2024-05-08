package source

import (
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/commands/source/add"
)

type Args struct {
	Add add.Args `command:"add" description:"Add a new source"`
}

func Run(args Args, subcommand *flags.Command) error {
	switch subcommand.Name {
	case "add":
		return add.Run(args.Add)
	}

	return nil
}
