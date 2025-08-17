package commands

import (
	"fmt"
	"context"
	"encoding/xml"
)

//Aggregator func
func HandlerAggregator(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("agg command requires no args")
	}

	rssURL := "https://www.wagslane.dev/index.xml"
	ctx := context.Background()

	feed, err := fetchFeed(ctx, rssURL)
	if err != nil {
		fmt.Printf("error getting rss: %s", err)
	}

	feed.Unescape()

	b, _ := xml.MarshalIndent(feed, "", "  ")
	fmt.Println(string(b))
	return nil
}

