package main

import (
	"fmt"
	"net/http"

	"github.com/BhanukarB/rssagg/internal/auth"
	"github.com/BhanukarB/rssagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("auth error %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("couldn't get user %v", err))
			return
		}

		handler(w, r, user)
	}
}
