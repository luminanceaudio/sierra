package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"sierra/services/sierra/commands/source"
)

var args struct {
	Source source.Args `command:"source" description:"Manage sources"`
}

func main() {
	p := flags.NewParser(&args, flags.Default)

	_, err := p.Parse()
	if err != nil {
		p.WriteHelp(os.Stderr)
		os.Exit(1)
	}

	switch p.Active.Name {
	case "source":
		err = source.Run(args.Source, p.Active.Active)
	}
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
