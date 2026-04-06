package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments> required")
	}

	ctx := context.Background()
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	feeds, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("error getting feed follows: %v", err)
	}

	fmt.Printf("feeds followed for user '%v':\n", user.Name)
	for _, feed := range feeds {
		fmt.Printf("  * %v\n", feed.FeedName)
	}

	return nil
}
