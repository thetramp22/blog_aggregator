package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments required")
	}

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error reseting database: %v", err)
	}
	fmt.Println("database reset")
	return nil
}
