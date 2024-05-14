package config

import (
	"fmt"
	"github.com/luminanceaudio/sierra/src/sierra/common/appdir"
	"os"
	"path/filepath"
)

func GetAppName() string {
	return "Sierra"
}

func CreateAppDataDir(extraPath string) (string, error) {
	appDataDir, err := appdir.GetApplicationDataDirForApp(GetAppName())
	if err != nil {
		return "", fmt.Errorf("could not get appdata dir: %s", err)
	}

	if extraPath != "" {
		appDataDir = filepath.Join(appDataDir, extraPath)
	}

	err = os.MkdirAll(appDataDir, 0750)
	if err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("failed to create app data directory '%s': %w", appDataDir, err)
	}

	return appDataDir, nil
}
