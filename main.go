package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/htet-29/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("no argument pass for login")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Printf("User %s is login\n", s.cfg.CurrentUsername)
	return nil
}

type commands struct {
	handlerFuncs map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.handlerFuncs[cmd.name]
	if !ok {
		return errors.New("the given handler is not registered")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlerFuncs[name] = f
}

func main() {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("not enough arguments were provided")
	}

	if args[1] == "login" && len(args) < 3 {
		log.Fatal("a username is required")
	}

	var s state
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Reading config: %+v\n", cfg)

	s.cfg = &cfg

	cmds := commands{
		handlerFuncs: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	cmd := command{
		name: args[1],
		args: []string{args[2]},
	}

	err = cmds.run(&s, cmd)
	if err != nil {
		log.Fatalf("cannot run command: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Reading config again: %+v\n", cfg)
}
