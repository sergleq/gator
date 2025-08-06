package cli

import (
	"context"
	"fmt"
	"gator/internal/state"
)

func HandlerListFollows(s *state.State, cmd Command) error {
	user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUser)
	if err != nil {
		return fmt.Errorf("пользователь не найден: %w", err)
	}

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("не удалось получить подписки: %w", err)
	}

	for _, f := range follows {
		fmt.Printf("- %s (%s) [by %s]\n", f.FeedName, f.FeedUrl, f.UserName)
	}

	return nil
}
