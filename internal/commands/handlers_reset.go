package commands

import (
	"fmt"
	"context"
)

//DANGER clears users database of ALL data
func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("reset command requires no args")
	}

	ctx := context.Background()
	err := s.Db.DeleteAllUsers(ctx)
	if err != nil {
		fmt.Printf("error deleting users: %s", err)
	}

	fmt.Println("Database has been reset.")
	return nil
}