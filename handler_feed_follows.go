package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/oleksandrdevhub/gorssagg/internal/database"
)

func (apiConfig *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feedFolow, err := apiConfig.DB.CreateFeedFollow(
		r.Context(),
		database.CreateFeedFollowParams{
			UserID:    user.ID,
			FeedID:    params.FeedID,
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
		},
	)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error creating feedFolow: %v", err))
		return
	}

	responseWithJSON(w, 201, databaseFeedFollowToFeedFollow(feedFolow))
}

func (apiConfig *apiConfig) handlerGetFeedFollowsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiConfig.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error getting feedFollows: %v", err))
		return
	}

	responseWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiConfig *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := chi.URLParam(r, "FeedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing feedFollowID: %v", err))
		return
	}

	err = apiConfig.DB.DeleteFeedFollow(
		r.Context(),
		database.DeleteFeedFollowParams{
			ID:     feedFollowID,
			UserID: user.ID,
		},
	)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error deleting feedFollow: %v", err))
		return
	}

	responseWithJSON(w, 204, struct{}{})
}
