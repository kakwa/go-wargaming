package wot


// AccountTanksOptions options.
type AccountTanksOptions struct {
	// Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
	AccessToken *string
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Localization language. Default is "ru". Valid values:
	// 
	// "en" - English 
	// "ru" - Русский (by default)
	// "pl" - Polski 
	// "de" - Deutsch 
	// "fr" - Français 
	// "es" - Español 
	// "zh-cn" - 简体中文 
	// "zh-tw" - 繁體中文 
	// "tr" - Türkçe 
	// "cs" - Čeština 
	// "th" - ไทย 
	// "vi" - Tiếng Việt 
	// "ko" - 한국어
	Language *string
	// Player's vehicle ID. Maximum limit: 100.
	TankId []int
}

type AccountTanks struct {
	// Mastery Badges:
	// 
	// 0 — None
	// 1 — 3rd Class 
	// 2 — 2nd Class
	// 3 — 1st Class
	// 4 — Ace Tanker
	MarkOfMastery *int `json:"mark_of_mastery,omitempty"`
	// Vehicle statistics
	Statistics *struct {
		// Battles fought
		Battles *int `json:"battles,omitempty"`
		// Victories
		Wins *int `json:"wins,omitempty"`
	} `json:"statistics,omitempty"`
	// Vehicle ID
	TankId *int `json:"tank_id,omitempty"`
}
