package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type UserPreference struct {
	db *sql.DB
}

func (r *UserPreference) Save(ctx context.Context, userID string, preferences interface{}) (err error) {
	var preferenceBytes []byte

	if preferenceBytes, err = json.Marshal(preferences); err != nil {
		return fmt.Errorf("preferences passed in an invalid format: %v", err)
	}

	if _, err = r.db.ExecContext(
		ctx,
		`
		INSERT INTO example.user_preference (user_id, preferences)
		VALUES ($1, $2::jsonb)
		ON CONFLICT (user_id)
		DO UPDATE SET preferences = excluded.preferences
		`,
		userID,
		preferenceBytes,
	); err != nil {
		log.Printf(
			"Could not save preferences for user %s: %v\n",
			userID, err,
		)
	}

	return err
}

func NewUserPreference(db *sql.DB) *UserPreference {
	return &UserPreference{
		db: db,
	}
}
