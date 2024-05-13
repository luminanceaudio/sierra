package server

import (
	"context"
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/cli/server/shutdown"
	"sierra/services/sierra/cli/server/start"
)

type Args struct {
	Start    start.Args    `command:"start" description:"Start HTTP server"`
	Shutdown shutdown.Args `command:"shutdown" description:"Shutdown running HTTP server"`
}

func Run(ctx context.Context, args Args, subcommand *flags.Command) error {
	switch subcommand.Name {
	case "start":
		return start.Run(args.Start)
	case "shutdown":
		return shutdown.Run(args.Shutdown)
	}

	return nil
}
