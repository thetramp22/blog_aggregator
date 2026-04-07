package main

import (
	"context"
	"fmt"

	"github.com/thetramp22/blog_aggregator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments> required")
	}

	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %v", err)
	}

	fmt.Printf("feeds followed for user '%v':\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("  * %v\n", feed.FeedName)
	}

	return nil
}
