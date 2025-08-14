package commands

import (
	"github.com/Piep220/go-blog-aggregator/internal/config"
	"github.com/Piep220/go-blog-aggregator/internal/database"
)

type State struct {
	Cfg  *config.Config
	Db   *database.Queries
}

type Command struct {
	Name    string
	Args    []string
}

type handlerFn func(s *State, cmd Command) error

type commands struct {
	m 		map[string]handlerFn
}