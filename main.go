package main

import (
	"fmt"
	"os"

	"github.com/abdielrumaldo/aggro-gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	curCommand, ok := c.commands[cmd.name]
	if !ok {
		return fmt.Errorf("Command '%s' not found\n", cmd.name)
	}
	return curCommand(s, cmd)
}

func handlerLogin(s *state, cmd command) error {

	if len(cmd.args) == 0 {
		return fmt.Errorf("No arguments provided for command '%s'", cmd.name)
	}
	// 0 is the username arg
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("user has been set!")
	return nil
}

func main() {
	cfg, err := config.Read()
	appState := state{
		cfg: &cfg,
	}
	if err != nil {
		println("Error reading config", err)
		os.Exit(1)
	}
	commands := commands{
		commands: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)

	fmt.Printf("Read config: %+v\n", cfg)

	cmdArgs := os.Args
	if len(cmdArgs) < 2 {
		fmt.Println("command name and args required")
		os.Exit(1)
	}
	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	if err := commands.run(&appState, cmd); err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	fmt.Printf("Read config again: %+v\n", cfg)
}
