package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments required")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds: %v", err)
	}

	for _, feed := range feeds {
		fmt.Printf("* %v\n", feed.FeedName)
		fmt.Printf("  url: %v\n", feed.Url)
		fmt.Printf("  added by user: %v\n", feed.UserName)
	}

	return nil
}
