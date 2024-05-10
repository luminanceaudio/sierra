package config

import (
	"fmt"
	"os"
	"sierra/common/appdir"
)

func GetAppName() string {
	return "Sierra"
}

// GetAppDataDir gets the application dir for Sierra
func GetAppDataDir() (string, error) {
	return appdir.GetApplicationDataDirForApp(GetAppName())
}

func CreateAppDataDir() error {
	appDataDir, err := GetAppDataDir()
	if err != nil {
		return fmt.Errorf("could not get appdata dir: %s", err)
	}

	err = os.Mkdir(appDataDir, 0755)
	if err != nil && !os.IsExist(err) {
		return fmt.Errorf("failed to create app data directory '%s': %w", appDataDir, err)
	}

	return nil
}
