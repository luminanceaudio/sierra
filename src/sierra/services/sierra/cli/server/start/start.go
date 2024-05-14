package start

import (
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/client"
	"github.com/luminanceaudio/sierra/src/sierra/services/sierra/internal/server"
)

type Args struct {
}

func Run(args Args) error {
	appClient := client.NewClient()

	// Best effort to shut down already running server
	_ = appClient.Shutdown()

	return server.StartBlocking()
}
