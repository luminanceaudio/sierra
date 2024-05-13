package pathfinder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetApplicationDataDirForApp gets the directory where app files can be saved
func GetApplicationDataDirForApp(appName string) (string, error) {
	// Windows
	appData := os.Getenv("APPDATA")
	if appData != "" {
		return filepath.Join(appData, appName), nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not get home dir: %s", err)
	}

	// NOTE: Due to sqlite problems with the dir having a space, it will not be used for now
	// MacOS
	//libraryDir := filepath.Join(homeDir, "Library", "Application Support")
	//_, err = os.Stat(libraryDir)
	//if err == nil {
	//	return filepath.Join(libraryDir, appName), nil
	//}

	// Linux
	return filepath.Join(homeDir, "."+strings.ToLower(appName)), nil
}
