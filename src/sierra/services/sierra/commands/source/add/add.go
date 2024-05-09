package add

import (
	"fmt"
	"sierra/services/sierra/internal/modules/configfile"
	"sierra/services/sierra/internal/modules/source"
)

type Args struct {
	Positional struct {
		Source string `positional-arg-name:"<source>" required:"true"`
	} `positional-args:"yes" required:"yes"`
}

func Run(args Args) error {
	config, err := configfile.Load()
	if err != nil {
		return fmt.Errorf("failed to get or create config: %s", err)
	}

	newSource, err := source.NewLocalSource(args.Positional.Source)
	if err != nil {
		return fmt.Errorf("failed to create source: %s", err)
	}

	config.Sources.Add(newSource)

	err = config.Save()
	if err != nil {
		return fmt.Errorf("failed to save config: %s", err)
	}

	return nil
}
