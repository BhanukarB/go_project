package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/BhanukarB/rssagg/internal/database"
	"github.com/google/uuid"
)

func(apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("error parsing JSON %s", err))
		return
	}

		now := time.Now().UTC()

	// Create a sql.NullTime for CreatedAt and UpdatedAt
	// Assuming your database.CreateUserParams expects sql.NullTime for these fields
	// and that your SQL table has DEFAULT CURRENT_TIMESTAMP, you might not
	// need to explicitly set UpdatedAt here if your DB handles it on insert.
	// However, if your DB query explicitly requires it, this is how you'd do it.
	createdAtNullTime := sql.NullTime{Time: now, Valid: true}
	updatedAtNullTime := sql.NullTime{Time: now, Valid: true}
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: createdAtNullTime,
		UpdatedAt: updatedAtNullTime,
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("couldn't create feed %v", err))
		return
	}
	respondWithJSON(w, 201, dbFeedToFeed(feed))
}

func(apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {

	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("couldn't get feeds %v", err))
		return
	}
	respondWithJSON(w, 200, dbFeedsToFeeds(feeds))
}

