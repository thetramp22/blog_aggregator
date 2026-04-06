package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments required")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %v", err)
	}
	for _, user := range users {
		line := fmt.Sprintf("* %v", user.Name)
		if user.Name == s.cfg.CurrentUserName {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
}
