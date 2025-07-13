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

func(apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
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
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: createdAtNullTime,
		UpdatedAt: updatedAtNullTime,
		Name: params.Name,
	})
	if err != nil{
		respondWithError(w, 400, fmt.Sprintf("couldn't create user %v", err))
		return
	}
	respondWithJSON(w, 201, dbUsertoUser(user))
}

func(apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	

	respondWithJSON(w, 200, dbUsertoUser(user))
}