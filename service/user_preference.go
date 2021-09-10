package service

import (
	"context"

	"github.com/testing/example/model"
)

type UserPreferenceRepository interface {
	Save(ctx context.Context, userID string, preferences model.UserPreferences) (err error)
}

type UserPreference struct {
	repo UserPreferenceRepository
}

func (u *UserPreference) Save(ctx context.Context, userID string, preferences model.UserPreferences) (err error) {
	return u.repo.Save(ctx, userID, preferences)
}

func NewUserPreference(repo UserPreferenceRepository) *UserPreference {
	return &UserPreference{
		repo: repo,
	}
}
