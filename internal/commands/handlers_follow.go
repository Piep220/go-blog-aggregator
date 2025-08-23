package commands

import (
	"context"
	"fmt"
	"net/url"
	"time"
	"github.com/google/uuid"
	"github.com/Piep220/go-blog-aggregator/internal/database"
)

//Add feed url to MiddlewareLoggedIn
func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("follow command requires one arg: url")
	}

	feedUrl := cmd.Args[0]
	//Check if URL is valid
	_, err := url.Parse(feedUrl)
	if err != nil {
		return fmt.Errorf("error parsing URL, please check, %w", err)
	}

	ctx := context.Background()
	feedID, err := s.Db.GetFeedFromUrl(ctx, feedUrl)
	if err != nil {
		return fmt.Errorf("error getting feed ID, %w", err)
	}

	nowTime := time.Now()
	newFeedFollow := database.CreateFeedFollowParams{
		ID: 		uuid.New(),
		CreatedAt: 	nowTime,
		UpdatedAt: 	nowTime,
		UserID: 	user.ID,
		FeedID: 	feedID.ID,	
	}

	_, err = s.Db.CreateFeedFollow(ctx, newFeedFollow)
	if err != nil {
		return fmt.Errorf("error CreateFeedFollow, %w", err)
	}

	fmt.Printf("\nCreated new feed-follow relation:\n")
	fmt.Printf("Feed: %s\nUser: %s\n", feedID.Name, user.Name)
	return nil
}

//Lists all follows from MiddlewareLoggedIn user
func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("follows command requires no args")
	}

	ctx := context.Background()
	follows, err := s.Db.GetFeedFollowsForUser(ctx, user.Name)
	if err != nil {
		return fmt.Errorf("error getting feeds1: %w", err)
	}

	if len(follows) == 0 {
		fmt.Printf("User: %s has no feedfollows configured.\n", user.Name)
		return nil
	}

	fmt.Printf("Current follows for %s:\n", user.Name)
	fmt.Printf("----- FOLLOW LIST -----\n")
	for _, follow := range follows {
		fmt.Printf("%s\n", follow.FeedName.String)
	}
	fmt.Printf("----- END OF LIST -----\n")
	return nil
}