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
	cmds.Register("addfeed", commands.MiddlewareLoggedIn(commands.HandlerAddFeed))
	cmds.Register("feeds", commands.HandlerListFeeds)
	cmds.Register("follow", commands.MiddlewareLoggedIn(commands.HandlerFollow))
	cmds.Register("following", commands.MiddlewareLoggedIn(commands.HandlerFollowing))
	cmds.Register("unfollow", commands.MiddlewareLoggedIn(commands.HandlerUnfollow))
	cmds.Register("browse", commands.MiddlewareLoggedIn(commands.HandlerBrowse))

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