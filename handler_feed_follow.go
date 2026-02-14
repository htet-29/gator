package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/htet-29/gator/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
)

func handlerCreateFeedFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.Name)
	}
	url := cmd.Args[0]
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	feed, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}

	user, err := s.db.GetUser(ctx, s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("couldn't get user: %w", err)
	}

	feedfollow, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	printFeedFollow(feedfollow)
	return nil
}

func handlerGetFeedFollows(s *state, cmd command) error {
	feeds, err := s.db.GetFeedFollowForUser(context.Background(), s.cfg.CurrentUsername)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}
	for _, feed := range feeds {
		printFeedFollowsForUser(feed)
	}
	return nil
}

func printFeedFollow(feed database.CreateFeedFollowRow) {
	fmt.Printf(" * Feed Name: %s\n", feed.FeedName)
	fmt.Printf(" * User Name: %s\n", feed.UserName)
}

func printFeedFollowsForUser(feed database.GetFeedFollowForUserRow) {
	fmt.Printf(" * Feed Name: %s | User Name: %s\n", feed.FeedName, feed.UserName)
}
