package cli

import (
	"context"
	"fmt"
	"gator/internal/state"
)

func HandlerReset(s *state.State, cmd Command) error {
	err := s.DB.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось очистить базу: %w", err)
	}
	fmt.Println("Все пользователи удалены.")
	return nil
}
