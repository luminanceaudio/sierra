package index

import (
	"context"
	"fmt"
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/internal/configfile"
	"sierra/services/sierra/internal/indexer"
)

type Args struct {
}

func Run(ctx context.Context, args Args, subcommand *flags.Command) error {
	config, err := configfile.Load()
	if err != nil {
		return err
	}

	idx := indexer.New()

	for _, oneofSource := range config.Sources.List {
		source, err := oneofSource.Get()
		if err != nil {
			return fmt.Errorf("failed getting source data: %s", err)
		}

		err = idx.Index(ctx, source)
		if err != nil {
			return fmt.Errorf("failed indexing source: %s", err)
		}
	}

	return nil
}
