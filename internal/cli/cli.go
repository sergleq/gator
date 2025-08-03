package cli

import (
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
)

type State struct {
	DB  *database.Queries
	CFG *config.Config
}
type Command struct {
	Name string
	Args []string
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
