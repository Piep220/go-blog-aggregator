package commands

import (
	"fmt"
)

//Runs handler function from command registery
func (c *commands) Run(s *State, cmd Command) error {
	h, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return h(s, cmd)
}

//Adds handlerFunctions to list of available commands
func (c *commands) Register(name string, f handlerFn) {
	c.registeredCommands[name] = f
}