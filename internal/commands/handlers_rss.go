package commands

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"strings"

	//"encoding/xml"
	"fmt"
	"net/url"
	"time"

	"github.com/Piep220/go-blog-aggregator/internal/database"
	"github.com/google/uuid"

	"github.com/araddon/dateparse"
)

//Use in loop, updates feeds by oldest first
func scrapeFeeds(s *State, ctx context.Context) error {
	//Get next feed from DB
	nextFeed, err := s.Db.GetNextFeedToFetch(ctx)
	if err != nil {
		return fmt.Errorf("error getting next feed: %s", err)
	}

	//Mark fetched
	err = s.Db.MarkFeedFetched(ctx, nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error marking fetched: %s", err)
	}

	//Fetch using fetchfeed
	feed, err := fetchFeed(ctx, nextFeed.Url)
	if err!= nil {
		return fmt.Errorf("error fetching url feed: %s", err)
	}
	/*
	//Iterate over feed, print titles
	for _, item := range feed.Channel.Item {
		fmt.Printf("%s\n", item.Title)
	}
	*/

	//Save feeds
	for _, item := range feed.Channel.Item {
		paresedTime, err := dateparse.ParseAny(item.PubDate)
		if err != nil {
			return fmt.Errorf("error parsing time in scrape feeds: %w", err)
		}
		post := database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt:  time.Now(),
    		UpdatedAt: time.Now(),
    		Title: item.Title,
    		Url: item.Link,
    		Description: sql.NullString{String: item.Description, Valid: item.Description != ""},
    		PublishedAt: paresedTime,
    		FeedID: nextFeed.ID,
		}
		_, err = s.Db.CreatePost(ctx, post)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			fmt.Printf("error scrapeFeed creating post: %s", err)
			continue
		}
	}
	return nil
}

//Brows posts, shows (default 2) posts
func HandlerBrowse(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("browse command requires zero or one arg: limit count")
	}

	browseLimit := 2
	if len(cmd.Args) == 1 {
		intArg, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("failure to convert arg to int: %w", err)
		}
		browseLimit = intArg
	}

	ctx := context.Background()
	getPostParams := database.GetPostsForUserByNameParams{
		Name: user.Name,
		Limit: int32(browseLimit),
	}
	results, err := s.Db.GetPostsForUserByName(ctx, getPostParams)
	if err != nil {
		return fmt.Errorf("error getting posts: %w", err)
	}

	fmt.Printf("Found %d posts for user %s:\n", len(results), user.Name)
	for _, post := range results {
		fmt.Printf("%s\n", post.PublishedAt.Format("Mon Jan 2"))
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}


	return nil
}

//Aggregator func
func HandlerAggregator(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("agg command requires one arg: time_between_req")
	}

	interval, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		fmt.Printf("error parsing time, string format: 1s, 2m, etc.")
	}

	ctx := context.Background()

	fmt.Printf("Collecting feeds every: %s\n", interval)
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		fmt.Printf("Running scraper\n")
		scrapeFeeds(s, ctx)
	}

	//feed.Unescape()

	//b, _ := xml.MarshalIndent(feed, "", "  ")
	//fmt.Println(string(b))

	//return nil
}

//Add feed by URL for MiddlewareLoggedIn user
func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("addfeed command requires two args, 'name' then 'url'")
	}

	//Check if URL is valid
	_, err := url.Parse(cmd.Args[1])
	if err != nil {
		return fmt.Errorf("error parsing URL, please check, %w", err)
	}

	ctx := context.Background()

	nowTime := time.Now()
	newFeedID := uuid.New()
	newFeed := database.AddFeedParams{
		ID:    	   newFeedID,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
		Name:	   cmd.Args[0],
		Url: 	   cmd.Args[1],
		UserID:    user.ID,
	}

	_, err = s.Db.AddFeed(ctx, newFeed)
	if err != nil {
		fmt.Printf("error adding feed: %s", err)
	}

	newFeedFollow := database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	nowTime,
		UpdatedAt: 	nowTime,
		UserID: 	user.ID,
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

//Lists all feeds configured
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
		fmt.Printf("Created by: %s\n\n", feed.UserName.String)
	}
	return nil
}