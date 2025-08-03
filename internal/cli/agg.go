package cli

import (
	"context"
	"fmt"
	"gator/internal/rssfeed"
)

func HandlerAgg(s *State, cmd Command) error {
	feed, err := rssfeed.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("ошибка загрузки ленты: %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}
