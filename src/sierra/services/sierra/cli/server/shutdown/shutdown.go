package shutdown

import "sierra/services/sierra/client"

type Args struct {
}

func Run(args Args) error {
	appClient := client.NewClient()
	aErr := appClient.Shutdown()
	if aErr != nil {
		return aErr
	}
	
	return nil
}
