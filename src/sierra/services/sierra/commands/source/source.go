package source

import (
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/commands/source/add"
	"sierra/services/sierra/commands/source/purge"
)

type Args struct {
	Add   add.Args   `command:"add" description:"Add a new source"`
	Purge purge.Args `command:"purge" description:"Remove all sources"`
}

func Run(args Args, subcommand *flags.Command) error {
	switch subcommand.Name {
	case "add":
		return add.Run(args.Add)
	case "purge":
		return purge.Run(args.Purge)
	}

	return nil
}
