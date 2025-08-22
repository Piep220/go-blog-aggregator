package commands

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/url"
	"time"
	"github.com/google/uuid"
	"github.com/Piep220/go-blog-aggregator/internal/database"
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

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed command requires two args, 'name' then 'url'")
	}

	//Check if URL is valid
	_, err := url.Parse(cmd.Args[1])
	if err != nil {
		return fmt.Errorf("error parsing URL, please check, %w", err)
	}

	ctx := context.Background()
	currentUser := s.Cfg.CurrentUserName
	userID, err := s.Db.GetUser(ctx, currentUser)
	if err != nil {
		return fmt.Errorf("error getting current user's ID, %w", err)
	}

	nowTime := time.Now()
	newFeedID := uuid.New()
	newFeed := database.AddFeedParams{
		ID:    	   newFeedID,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
		Name:	   cmd.Args[0],
		Url: 	   cmd.Args[1],
		UserID:    userID.ID,
	}

	_, err = s.Db.AddFeed(ctx, newFeed)
	if err != nil {
		fmt.Printf("error adding feed: %s", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	nowTime,
		UpdatedAt: 	nowTime,
		UserID: 	userID.ID,
		FeedID: 	newFeedID,	
	}

	_, err = s.Db.CreateFeedFollow(ctx, newFeedFollow)
	if err != nil {
		return fmt.Errorf("error CreateFeedFollow, %w", err)
	}


	fmt.Printf("Feed: %s, created.\n", cmd.Args[0])
	if b, err := json.MarshalIndent(newFeed, "", "  "); err == nil {
        fmt.Println(string(b))
    }


	return nil
}

func HandlerListFeeds(s *State, cmd Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("feeds command requires no args")
	}

	ctx := context.Background()
	feeds, err := s.Db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("error getting feeds: %w", err)
	}

	if len(feeds) == 0 {
		fmt.Println("You have no feeds configured.")
		return nil
	}

	fmt.Println("Current feeds are: ")
	for _, feed := range feeds {
		fmt.Printf("Name: %s\n", feed.Name)
		fmt.Printf("URL:  %s\n", feed.Url)
		fmt.Printf("Created by: %s\n\n", feed.Name_2.String)
	}
	return nil
}