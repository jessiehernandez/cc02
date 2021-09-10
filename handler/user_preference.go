package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type postBody struct {
	Preferences map[string]interface{} `json:"preferences"`
	UserID      string                 `json:"userID"`
}

type UserPreferenceService interface {
	Save(ctx context.Context, userID string, preferences interface{}) (err error)
}

type UserPreference struct {
	userPreferenceService UserPreferenceService
}

func (u *UserPreference) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	ctx := r.Context()

	if r.Method == http.MethodPost {
		requestBody := postBody{}

		if err = json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Invalid JSON body"))
			return
		}

		if requestBody.UserID == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("User ID not specified"))
			return
		}

		if err = u.userPreferenceService.Save(ctx, requestBody.UserID, requestBody.Preferences); err != nil {
			log.Printf("Could not save user preferences: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("Could not save user preferences"))
			return
		}

		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewUserPreference(userPreferenceService UserPreferenceService) *UserPreference {
	return &UserPreference{
		userPreferenceService: userPreferenceService,
	}
}
