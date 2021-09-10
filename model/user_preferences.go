package model

type UserPreferences struct {
	Locale   string `json:"locale"`
	Timezone string `json:"timezone"`
}
