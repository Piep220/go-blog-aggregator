package config

import (
	"encoding/json"
	"os"
	"fmt"
)

func Read() (*Config, error) {
	configFilePath, err := getConfigFilePath()
	if err != nil {
		fmt.Printf("cannot get config file path: %v\n", err)
		return nil, err
	}

	data, err := os.ReadFile(configFilePath)
    if err != nil {
        fmt.Printf("cannot read file: %v\n", err)
		return nil, err
	}

	conf := &Config{}
	err = json.Unmarshal(data, conf)
	if err != nil {
		fmt.Printf("cannot unmarshal config: %v\n", err)
		return nil, err
	}

	return conf, nil
}