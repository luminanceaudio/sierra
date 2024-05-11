package add

import (
	"context"
	"sierra/services/sierra/internal/modules/source"
)

type Args struct {
	Positional struct {
		Path string `positional-arg-name:"<path>" required:"true"`
	} `positional-args:"yes" required:"yes"`
}

func Run(ctx context.Context, args Args) error {
	err := source.CreateLocal(ctx, args.Positional.Path)
	if err != nil {
		return err
	}

	return nil
}
