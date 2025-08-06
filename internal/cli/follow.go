package cli

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"
	"gator/internal/state"

	"github.com/google/uuid"
)

func HandlerFollow(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("использование: follow <feed_url>")
	}

	url := cmd.Args[0]

	// Получаем текущего пользователя
	user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUser)
	if err != nil {
		return fmt.Errorf("не удалось получить текущего пользователя: %w", err)
	}

	// Ищем фид по URL
	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("фид с таким URL не найден: %w", err)
	}

	// Создаем запись о подписке
	params := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.DB.CreateFeedFollow(context.Background(), params)
	if err != nil {
		return fmt.Errorf("не удалось создать подписку: %w", err)
	}

	fmt.Printf("Пользователь %s подписался на фид %s\n", follow.UserName, follow.FeedName)
	return nil
}
