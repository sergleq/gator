package agg

import (
	"context"
	"fmt"
	"gator/internal/rssfeed"
	"gator/internal/state"
)

func ScrapeFeeds(s *state.State) {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		fmt.Println("Нет доступных фидов:", err)
		return
	}

	err = s.DB.MarkFeedFetched(context.Background(), feed.ID)
	if err != nil {
		fmt.Println("Ошибка при обновлении времени фида:", err)
		return
	}

	rss, err := rssfeed.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		fmt.Printf("Ошибка при парсинге фида %s: %v\n", feed.Url, err)
		return
	}

	fmt.Printf("📥 Фид: %s\n", feed.Name)
	for _, item := range rss.Channel.Items {
		fmt.Printf("- %s\n", item.Title)
	}
}
