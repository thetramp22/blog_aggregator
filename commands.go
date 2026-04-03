package main

import (
	"fmt"
)

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

type command struct {
	name string
	args []string
}

func (c commands) run(s *state, cmd command) error {
	if _, exists := c.registeredCommands[cmd.name]; !exists {
		return fmt.Errorf("command '%v' not valid", cmd.name)
	}
	err := c.registeredCommands[cmd.name](s, cmd)
	if err != nil {
		return err
	}
	return nil
}

func (c commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
