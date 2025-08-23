package commands

import (
	"github.com/Piep220/go-blog-aggregator/internal/config"
	"github.com/Piep220/go-blog-aggregator/internal/database"
)

//Program's internal state struct
type State struct {
	Cfg  *config.Config
	Db   *database.Queries
}

type Command struct {
	Name    string
	Args    []string
}

//Template for handlerFunctions
type handlerFn func(s *State, cmd Command) error
type handlerLoggedIn func(s *State, cmd Command, user database.User) error

//List of commands available to run
type commands struct {
	registeredCommands 		map[string]handlerFn
}