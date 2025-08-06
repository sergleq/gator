package cli

import (
	"context"
	"fmt"
)

func HandlerFollowing(s *State, cmd Command) error {
	// Получаем текущего пользователя
	user, err := s.DB.GetUser(context.Background(), s.CFG.CurrentUser)
	if err != nil {
		return fmt.Errorf("пользователь не найден: %w", err)
	}

	// Получаем список подписок
	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("не удалось получить подписки: %w", err)
	}

	if len(follows) == 0 {
		fmt.Println("Вы еще не подписаны ни на один фид.")
		return nil
	}

	fmt.Println("Подписки пользователя", user.Name, ":")
	for _, f := range follows {
		fmt.Printf("- %s\n", f.FeedName)
	}

	return nil
}
