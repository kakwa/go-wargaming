package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
)

type WgnAccountInfo struct {
	// Player ID
	AccountId int `json:"account_id,omitempty"`
	// Date when player's account was created
	CreatedAt UnixTime `json:"created_at,omitempty"`
	// List of games player has played
	Games []string `json:"games,omitempty"`
	// Player name
	Nickname string `json:"nickname,omitempty"`
	// Player's private data
	Private struct {
		// Amount of Free Experience
		FreeXp int `json:"free_xp,omitempty"`
		// Current gold balance
		Gold int `json:"gold,omitempty"`
		// Premium Account expiration date
		PremiumExpiresAt UnixTime `json:"premium_expires_at,omitempty"`
	} `json:"private,omitempty"`
}

// WgnAccountInfo Method returns Wargaming account details.
//
// https://developers.wargaming.net/reference/all/wgn/account/info
//
// account_id:
//     Player ID. Maximum limit: 100.
// access_token:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
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
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "cs" &mdash; Čeština 
//     "th" &mdash; ไทย 
//     "vi" &mdash; Tiếng Việt 
//     "ko" &mdash; 한국어
func (client *Client) WgnAccountInfo(realm Realm, accountId []int, accessToken string, fields []string, language string) (*WgnAccountInfo, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"access_token": accessToken,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *WgnAccountInfo
	err := client.doGetDataRequest(realm, "/wgn/account/info/", reqParam, &result)
	return result, err
}
