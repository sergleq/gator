package cli

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("использование: unfollow <feed_url>")
	}

	url := cmd.Args[0]

	params := database.DeleteFeedFollowByUserAndURLParams{
		UserID: user.ID,
		Url:    url,
	}

	err := s.DB.DeleteFeedFollowByUserAndURL(context.Background(), params)
	if err != nil {
		return fmt.Errorf("не удалось отписаться: %w", err)
	}

	fmt.Printf("Пользователь %s отписался от фида: %s\n", user.Name, url)
	return nil
}
