package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("no arguments required")
	}

	feedURL := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("error fetching '%v': %v", feedURL, err)
	}
	fmt.Printf("%+v\n", feed)
	return nil
}
