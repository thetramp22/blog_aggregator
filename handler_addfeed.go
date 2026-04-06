package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thetramp22/blog_aggregator/internal/database"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return fmt.Errorf("2 arguments required: <name> <url>")
	}
	name := cmd.args[0]
	url := cmd.args[1]
	ctx := context.Background()

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}

	newFeed, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("feed already exists: %v", err)
	}
	fmt.Printf("feed '%v' was added\n", newFeed.Name)
	fmt.Printf("%+v\n", newFeed)

	return nil
}
