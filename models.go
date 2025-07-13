package main

import (
	"database/sql"

	"github.com/BhanukarB/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	APIKey    string        `json:"api_key"`
}

func dbUsertoUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		APIKey: dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID     `json:"id"`
	Name      string        `json:"name"`
	CreatedAt sql.NullTime  `json:"created_at"`
	UpdatedAt sql.NullTime  `json:"updated_at"`
	URL       string        `json:"url"`
	UserID    uuid.UUID     `json:"user_id"`
}

func dbFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Name: dbFeed.Name,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		URL: dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func dbFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _,dbFeed := range dbFeeds {
		feeds = append(feeds, dbFeedToFeed(dbFeed))
	}
	return feeds
}