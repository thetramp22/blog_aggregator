package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("1 argument required: <time_between_reqs>")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("error parsing duration: %v", err)
	}

	fmt.Printf("collecting feeds every %v\n", timeBetweenRequests)
	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}

	// feedURL := "https://www.wagslane.dev/index.xml"
	// feed, err := fetchFeed(context.Background(), feedURL)
	// if err != nil {
	// 	return fmt.Errorf("error fetching '%v': %v", feedURL, err)
	// }
	// fmt.Printf("%+v\n", feed)
	// return nil
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed to scrape: %v", err)
	}

	s.db.MarkFeedFetched(context.Background(), nextFeed.ID)

	fetchedFeed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error getting feed: %v", err)
	}

	fmt.Println("===============")
	fmt.Printf("* %v:\n", fetchedFeed.Channel.Title)
	for _, item := range fetchedFeed.Channel.Item {
		fmt.Printf(" - %v\n", item.Title)
	}
	fmt.Println("===============")

	return nil
}
