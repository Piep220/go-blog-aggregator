package commands

import (
	"fmt"
)

func (c *commands) Run(s *State, cmd Command) error {
	h, ok := c.m[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return h(s, cmd)
}

func (c *commands) Register(name string, f handlerFn) {
	c.m[name] = f
}