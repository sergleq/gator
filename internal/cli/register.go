package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
	"time"

	"github.com/google/uuid"
)

func HandlerRegister(s *State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("регистрация требует имя пользователя")
	}

	username := cmd.Args[0]

	// создаём пользователя
	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      username,
	}

	_, err := s.DB.CreateUser(context.Background(), newUser)
	if err != nil {
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	// обновляем конфиг
	err = s.CFG.SetUser(username)
	if err != nil {
		return fmt.Errorf("ошибка обновления конфигурации: %w", err)
	}

	// лог и подтверждение
	fmt.Printf("Пользователь %q создан и установлен текущим\n", username)
	fmt.Printf("DEBUG: %+v\n", newUser)

	return nil
}
