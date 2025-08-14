package config

import (
	"fmt"
)

func (c *Config)SetUser(name string) error {
	c.CurrentUserName = name
	err := write(c)
	if err != nil {
		fmt.Printf("cannot write config after setting user: %v\n", err)
		return err
	}
	return nil
}