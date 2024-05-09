package configfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sierra/services/sierra/internal/modules/sources"
)

const (
	sierraUserConfigDirName = "Sierra"
	configFileName          = "config.json"
)

type Config struct {
	Sources *sources.Sources `json:"sources"`
}

func newConfig() *Config {
	return &Config{
		Sources: sources.New(),
	}
}

// Load loads the config file from disk or creates a new one
func Load() (*Config, error) {
	configFilePath, err := getSierraConfigFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get sierra config directory: %w", err)
	}

	// Create new config file if doesn't exist
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		c := newConfig()
		return c, c.Save()
	}

	configFile, err := os.OpenFile(configFilePath, os.O_RDWR, 0600)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	var config Config
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

func (c *Config) Save() error {
	configPath, err := getSierraConfigFilePath()
	if err != nil {
		return fmt.Errorf("failed to get sierra config directory: %w", err)
	}

	err = os.MkdirAll(filepath.Dir(configPath), 0700)
	if err != nil {
		return fmt.Errorf("failed to create directory %s: %w", filepath.Dir(configPath), err)
	}

	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config to file: %w", err)
	}

	return nil
}

func getSierraConfigFilePath() (string, error) {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return "", fmt.Errorf("failed getting user config dir: %w", err)
	}

	return filepath.Join(userConfigDir, sierraUserConfigDirName, configFileName), nil
}
