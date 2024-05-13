package source

import (
	"context"
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/cli/source/add"
)

type Args struct {
	Add add.Args `command:"add" description:"Add a new source"`
}

func Run(ctx context.Context, args Args, subcommand *flags.Command) error {
	switch subcommand.Name {
	case "add":
		return add.Run(ctx, args.Add)
	}

	return nil
}
