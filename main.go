package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thetramp22/blog_aggregator/internal/config"
)

type state struct {
	config *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v", err)
	}

	programState := &state{
		config: &cfg,
	}

	cmds := commands{
		registeredCommands: map[string]func(*state, command) error{},
	}
	cmds.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("command argument required")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]

	cmd := command{name: cmdName, args: cmdArgs}

	err = cmds.run(programState, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", programState.config)
}
