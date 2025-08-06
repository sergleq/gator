package cli

import (
	"context"
	"fmt"
	"time"

	"gator/internal/database"

	"github.com/google/uuid"
)

func MiddlewareLoggedIn(handler func(s *State, cmd Command, user database.User) error) func(*State, Command) error {
	return func(s *State, cmd Command) error {
		user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUser)
		if err != nil {
			return fmt.Errorf("требуется залогиненный пользователь: %w", err)
		}
		return handler(s, cmd, user)
	}
}

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("использование: addfeed <name> <url>")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	now := time.Now()
	feedID := uuid.New()

	feedParams := database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.DB.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("не удалось создать фид: %w", err)
	}

	fmt.Printf("Фид создан: %s (%s)\n", feed.Name, feed.Url)

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	follow, err := s.DB.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("не удалось создать подписку на фид: %w", err)
	}

	fmt.Printf("Подписка создана для пользователя %s на фид %s\n", follow.UserName, follow.FeedName)
	return nil
}
