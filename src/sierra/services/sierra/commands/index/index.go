package index

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/internal/modules/configfile"
	"sierra/services/sierra/internal/modules/indexer"
)

type Args struct {
}

func Run(args Args, subcommand *flags.Command) error {
	config, err := configfile.Load()
	if err != nil {
		return err
	}

	for _, oneofSource := range config.Sources.List {
		source, err := oneofSource.Get()
		if err != nil {
			return fmt.Errorf("failed getting source data: %s", err)
		}

		idx, err := indexer.New(source)
		if err != nil {
			return fmt.Errorf("failed creating indexer: %s", err)
		}

		_ = idx

		entries, err := source.ReadDir("")
		if err != nil {
			return fmt.Errorf("failed reading directory: %s", err)
		}

		for _, entry := range entries {
			fmt.Println(entry.Name())
		}
	}

	return nil
}
