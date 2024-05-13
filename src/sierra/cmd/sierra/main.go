package main

import (
	"context"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"sierra/services/sierra/cli/index"
	"sierra/services/sierra/cli/server"
	"sierra/services/sierra/cli/source"
)

var args struct {
	Source source.Args `command:"source" description:"Manage sources"`
	Index  index.Args  `command:"index" description:"Indexes sources"`
	Server server.Args `command:"server" description:"Run HTTP server"`
}

func main() {
	p := flags.NewParser(&args, flags.Default)

	_, err := p.Parse()
	if err != nil {
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	ctx := context.Background()

	switch p.Active.Name {
	case "source":
		err = source.Run(ctx, args.Source, p.Active.Active)
	case "index":
		err = index.Run(ctx, args.Index, p.Active.Active)
	case "server":
		err = server.Run(ctx, args.Server, p.Active.Active)
	}
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
