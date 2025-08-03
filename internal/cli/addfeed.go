package cli

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("использование: addfeed <name> <url>")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	// Получаем текущего пользователя
	user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUser)
	if err != nil {
		return fmt.Errorf("ошибка получения текущего пользователя: %w", err)
	}

	// Создаём фид
	params := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), params)
	if err != nil {
		return fmt.Errorf("ошибка создания фида: %w", err)
	}

	fmt.Println("Фид успешно создан:")
	fmt.Printf("  ID: %s\n  Name: %s\n  URL: %s\n  UserID: %s\n", feed.ID, feed.Name, feed.Url, feed.UserID)
	return nil
}
