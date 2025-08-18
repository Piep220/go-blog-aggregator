package main

import (
	"fmt"
	"os"
	"github.com/Piep220/go-blog-aggregator/internal/commands"
	"github.com/Piep220/go-blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

/*
Need to goose postgress down/up for submits.
Config in ~/.gatorconfig.json
*/


//Moved error handelling etc to main_helpers, to clarify.

func main() {
	fmt.Println("Starting...")
	cfg := readConfig()
	cmd := getCmd()

	cmds := commands.NewCommands()
	cmds.Register("login", commands.HandlerLogin)
	cmds.Register("register", commands.HandlerRegisterUser)
	cmds.Register("reset", commands.HandlerReset)
	cmds.Register("users", commands.HandlerPrintUsers)
	cmds.Register("agg", commands.HandlerAggregator)
	cmds.Register("addfeed", commands.HandlerAddFeed)
	cmds.Register("feeds", commands.HandlerListFeeds)

	db := dbOpen(cfg)
	defer db.Close()

	dbQueries := database.New(db)

	programState := &commands.State{
		Cfg: cfg,
		Db:  dbQueries,
	}

	err := cmds.Run(programState, cmd)
	if err != nil {
		fmt.Printf("error running command: %s", err)
		os.Exit(1)
	}
}