package main

import (
	"context"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"sierra/services/sierra/commands/index"
	"sierra/services/sierra/commands/source"
	"sierra/services/sierra/config"
)

var args struct {
	Source source.Args `command:"source" description:"Manage sources"`
	Index  index.Args  `command:"index" description:"Indexes sources"`
}

func main() {
	p := flags.NewParser(&args, flags.Default)

	_, err := p.Parse()
	if err != nil {
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	ctx := context.Background()

	err = config.CreateAppDataDir()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch p.Active.Name {
	case "source":
		err = source.Run(ctx, args.Source, p.Active.Active)
	case "index":
		err = index.Run(ctx, args.Index, p.Active.Active)
	}
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
