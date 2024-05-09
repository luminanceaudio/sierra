package purge

import (
	"fmt"
	"sierra/services/sierra/internal/modules/configfile"
)

type Args struct {
}

func Run(args Args) error {
	config, err := configfile.Load()
	if err != nil {
		return fmt.Errorf("failed to get or create config: %s", err)
	}

	config.Sources.Purge()

	err = config.Save()
	if err != nil {
		return fmt.Errorf("failed to save config: %s", err)
	}

	return nil
}
