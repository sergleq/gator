package cli

import (
	"context"
	"errors"
	"fmt"
	"gator/internal/state"
)

func HandlerLogin(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return errors.New("the login handler expects a single argument, the username")
	}
	username := cmd.Args[0]
	// Проверяем, существует ли пользователь в БД
	_, err := s.DB.GetUser(context.Background(), username)
	if err != nil {
		// Ошибка — пользователь не найден
		return fmt.Errorf("пользователь %q не найден", username)
	}

	// Если найден — записываем в конфиг
	err = s.CFG.SetUser(username)
	if err != nil {
		return fmt.Errorf("ошибка записи в конфиг: %w", err)
	}

	fmt.Println("User has been set")
	return nil
}
