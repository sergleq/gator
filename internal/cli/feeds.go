package cli

import (
	"context"
	"fmt"
	"gator/internal/state"
)

func HandlerFeeds(s *state.State, cmd Command) error {
	feeds, err := s.DB.GetFeedsWithUsers(context.Background())
	if err != nil {
		return fmt.Errorf("не удалось получить фиды: %w", err)
	}

	for _, f := range feeds {
		fmt.Printf("- %s (%s) [by %s]\n", f.FeedName, f.FeedUrl, f.UserName)
	}

	return nil
}
