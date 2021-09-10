package service

import "context"

type UserPreferenceRepository interface {
	Save(ctx context.Context, userID string, preferences interface{}) (err error)
}

type UserPreference struct {
	repo UserPreferenceRepository
}

func (u *UserPreference) Save(ctx context.Context, userID string, preferences interface{}) (err error) {
	return u.repo.Save(ctx, userID, preferences)
}

func NewUserPreference(repo UserPreferenceRepository) *UserPreference {
	return &UserPreference{
		repo: repo,
	}
}
