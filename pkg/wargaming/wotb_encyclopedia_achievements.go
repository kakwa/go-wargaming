package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strings"
)

// WotbEncyclopediaAchievements Method returns information about achievements.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/achievements
//
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
func (client *Client) WotbEncyclopediaAchievements(realm Realm, fields []string, language string) (*wotb.EncyclopediaAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotb.EncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
