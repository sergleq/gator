package cli

import (
	"errors"
	"fmt"
	"gator/internal/config"
)

type State struct {
	ConfigFile *config.Config
}
type Command struct {
	Name     string
	Argument []string
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Argument) < 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.Argument[0]
	s.ConfigFile.SetUser(username)
	fmt.Println("User has been set")
	return nil
}

type Commands struct {
	Handlers map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	handler, exists := c.Handlers[cmd.Name]
	if !exists {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}
	return handler(s, cmd)
}
func (c *Commands) Register(name string, f func(*State, Command) error) {
	c.Handlers[name] = f
}
