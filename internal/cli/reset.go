package cli

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось очистить базу: %w", err)
	}
	fmt.Println("Все пользователи удалены.")
	return nil
}
