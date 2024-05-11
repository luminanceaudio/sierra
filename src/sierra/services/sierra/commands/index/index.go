package index

import (
	"context"
	"fmt"
	"github.com/jessevdk/go-flags"
	"sierra/services/sierra/internal/indexer"
	"sierra/services/sierra/internal/modules/source"
)

type Args struct {
	Force bool `short:"f" long:"force" description:"force reindex"`
}

func Run(ctx context.Context, args Args, subcommand *flags.Command) error {
	sources, err := source.GetAll(ctx)
	if err != nil {
		return err
	}

	idx := indexer.New()

	for _, src := range sources {
		err = idx.Index(ctx, src, args.Force)
		if err != nil {
			return fmt.Errorf("failed indexing source: %s", err)
		}
	}

	return nil
}
