package cli

import (
	"fmt"
	"gator/internal/agg"
	"gator/internal/state"
	"time"
)

func HandlerAgg(s *state.State, cmd Command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("использование: agg <time_between_reqs>, например: agg 30s")
	}

	intervalStr := cmd.Args[0]
	dur, err := time.ParseDuration(intervalStr)
	if err != nil {
		return fmt.Errorf("некорректный интервал: %w", err)
	}

	fmt.Printf("Запущена агрегация фидов каждые %s\n", dur)

	ticker := time.NewTicker(dur)
	defer ticker.Stop()

	// первый запуск сразу
	agg.ScrapeFeeds(s)

	for range ticker.C {
		agg.ScrapeFeeds(s)
	}

	return nil
}
