package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("1 argument <username> required")
	}
	username := cmd.args[0]
	err := s.config.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("current user set to: %v\n", s.config.CurrentUserName)
	return nil
}
