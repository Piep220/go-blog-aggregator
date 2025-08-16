package commands

import (
	"context"
	"fmt"
	"os"
	"time"
	"encoding/json"
	"github.com/Piep220/go-blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login command requires only a username")
	}

	userName := cmd.Args[0]

	ctx := context.Background()
	_, err := s.Db.GetUser(ctx, userName)
	if err != nil {
		fmt.Printf("user not found, with error: %s\n", err)
		os.Exit(1)
	}

	err = s.Cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Printf("User has been set to: %q\n", userName)
	return nil
}

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("login command requires only a username")
	}

	nowTime := time.Now()
	ctx := context.Background()
	newUser := database.CreateUserParams{
		ID:    	   uuid.New(),
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
		Name:	   cmd.Args[0],
	}

	_, err := s.Db.CreateUser(ctx, newUser)
	if err != nil {
		fmt.Printf("error adding user: %s", err)
		os.Exit(1)
	}

	s.Cfg.SetUser(cmd.Args[0])
	fmt.Printf("User: %s, created.\n", cmd.Args[0])
	if b, err := json.MarshalIndent(newUser, "", "  "); err == nil {
        fmt.Println(string(b))
    }

	return nil
}

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command requires no args")
	}

	ctx := context.Background()
	err := s.Db.DeleteAllUsers(ctx)
	if err != nil {
		fmt.Printf("error deleting users: %s", err)
		os.Exit(1)
	}

	fmt.Println("Database has been reset.")
	return nil
}

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("users command requires no args")
	}

	ctx := context.Background()
	users, err := s.Db.GetUsers(ctx)
	if err != nil {
		fmt.Printf("error getting users: %s", err)
		os.Exit(1)
	}

	for _, user := range users {
		if s.Cfg.CurrentUserName == user {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}
	return nil
}