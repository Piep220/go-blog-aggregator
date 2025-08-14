package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get user home directory: %v", err)
	}

	configFilePath := dir + "/" + configFileName
	return configFilePath, nil
}

func write(conf *Config) error {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("cannot get config file path: %v\n", err)
		return err
	}

	data, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		fmt.Printf("cannot marshal config: %v\n", err)
		return err
	}

	err = os.WriteFile(configFilePath, data, 0644)
	if err != nil {
		fmt.Printf("cannot write file: %v\n", err)
		return err
	}

	return nil
}
