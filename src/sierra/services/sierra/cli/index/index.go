package index

import (
	"context"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/indexer"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/modules/source"
)

type Args struct {
	Force bool `short:"f" long:"force" description:"force reindex"`
}

func Run(ctx context.Context, args Args, subcommand *flags.Command) error {
	sources, err := source.GetAll(ctx)
	if err != nil {
		return err
	}

	for _, src := range sources {
		err = indexer.Singleton().Index(ctx, src, args.Force, false)
		if err != nil {
			return fmt.Errorf("failed indexing source: %s", err)
		}
	}

	return nil
}
