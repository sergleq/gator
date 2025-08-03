package cli

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	users, err := s.DB.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось получить пользователей: %w", err)
	}
	for _, u := range users {
		line := "* " + u.Name
		if u.Name == s.CFG.CurrentUser {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
}
