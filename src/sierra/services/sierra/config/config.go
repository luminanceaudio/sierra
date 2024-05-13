package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sierra/common/pathfinder"
)

func GetAppName() string {
	return "Sierra"
}

// GetAppDataDir gets the application dir for Sierra
func GetAppDataDir() (string, error) {
	return pathfinder.GetApplicationDataDirForApp(GetAppName())
}

func CreateAppDataDir() (string, error) {
	return createAppDataDirEx("")
}

func CreateAppImageCacheDir() (string, error) {
	return createAppDataDirEx("cache/images")
}

func createAppDataDirEx(innerDirs string) (string, error) {
	dir, err := GetAppDataDir()
	if err != nil {
		return "", fmt.Errorf("could not get appdata dir: %s", err)
	}

	if innerDirs != "" {
		dir = filepath.Join(dir, innerDirs)
	}

	err = os.MkdirAll(dir, 0755)
	if err != nil && !os.IsExist(err) {
		return "", fmt.Errorf("failed to create directory '%s': %w", dir, err)
	}

	return dir, nil
}
