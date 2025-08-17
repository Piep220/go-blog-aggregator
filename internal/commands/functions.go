package commands

func NewCommands() *commands {
	return &commands{registeredCommands: make(map[string]handlerFn)}
}