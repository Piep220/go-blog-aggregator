package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Piep220/go-blog-aggregator/internal/commands"
	"github.com/Piep220/go-blog-aggregator/internal/config"
)

//b, _ := json.MarshalIndent(cfg, "", "  ")
//fmt.Println(string(b))

func readConfig() *config.Config {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		return nil
	}
	return cfg
}

func getCmd() commands.Command {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Error: arguemnt required, none given")
		fmt.Println("Usage: mycli <command> [args...]")
		os.Exit(1)
	}
	
	//0th arg is program name
	//1th arg is command 
	//2th + arg, args for command
	cmd := commands.Command{
		Name: args[1],
		Args: args[2:],
	}

	return cmd
}

func dbOpen(cfg *config.Config) *sql.DB {
	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Printf("error opening databse: %s", err)
		os.Exit(1)
	}
	return db
}