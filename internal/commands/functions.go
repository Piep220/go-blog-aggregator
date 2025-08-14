package commands

func NewCommands() *commands {
	return &commands{m: make(map[string]handlerFn)}
}