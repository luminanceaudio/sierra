package add

import (
	"context"
	"sierra/services/sierra/internal/modules/source"
)

type Args struct {
	Positional struct {
		URI string `positional-arg-name:"<uri>" required:"true"`
	} `positional-args:"yes" required:"yes"`
}

func Run(ctx context.Context, args Args) error {
	err := source.Create(ctx, args.Positional.URI)
	if err != nil {
		return err
	}

	return nil
}
