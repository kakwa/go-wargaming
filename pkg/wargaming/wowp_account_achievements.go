package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WowpAccountAchievements struct {
	// Player account ID
	AccountId int `json:"account_id,omitempty"`
	// Achievements earned
	Achievements struct {
		// Number
		Count int `json:"count,omitempty"`
		// First time when an achievement was received
		FirstAt UnixTime `json:"first_at,omitempty"`
		// Last time when an achievement was received
		LastAt UnixTime `json:"last_at,omitempty"`
		// Most number
		MaxCount int `json:"max_count,omitempty"`
	} `json:"achievements,omitempty"`
}

// WowpAccountAchievements Method returns players' achievement details.
//
// https://developers.wargaming.net/reference/all/wowp/account/achievements
//
// account_id:
//     Player account ID. Maximum limit: 100.
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "en" &mdash; English 
//     "ru" &mdash; Русский (by default)
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "zh-cn" &mdash; 简体中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WowpAccountAchievements(realm Realm, accountId []int, fields []string, language string) (*WowpAccountAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WowpAccountAchievements
	err := client.doGetDataRequest(realm, "/wowp/account/achievements/", reqParam, &result)
	return result, err
}
