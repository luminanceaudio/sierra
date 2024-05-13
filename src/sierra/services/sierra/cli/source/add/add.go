package add

import (
	"context"
	"sierra/common/uri"
	"sierra/services/sierra/internal/modules/source"
)

type Args struct {
	Positional struct {
		URI string `positional-arg-name:"<uri>" required:"true"`
	} `positional-args:"yes" required:"yes"`
}

func Run(ctx context.Context, args Args) error {
	uri, err := uri.New(args.Positional.URI)
	if err != nil {
		return err
	}

	err = source.Create(ctx, uri)
	if err != nil {
		return err
	}

	return nil
}
