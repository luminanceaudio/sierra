package server

import (
	"context"
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/internal/server"
)

type Args struct {
}

func Run(ctx context.Context, args Args, subcommand *flags.Command) error {
	server.StartBlocking()
	return nil
}
