package config

import (
	"encoding/json"
	"os"
	"fmt"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	dir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot get user home directory: %v", err)
	}

	configFilePath := dir + "/" + configFileName
	return configFilePath, nil
}

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

func (c *Config)SetUser(name string) error {
	c.CurrentUserName = name
	err := write(c)
	if err != nil {
		fmt.Printf("cannot write config after setting user: %v\n", err)
		return err
	}
	return nil
}