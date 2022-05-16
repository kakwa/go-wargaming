package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotx"
	"strings"
)

// WotxEncyclopediaAchievements Method returns list of awards, medals, ribbons.
//
// https://developers.wargaming.net/reference/all/wotx/encyclopedia/achievements
//
// category:
//     Filter by award category. Maximum limit: 100. Valid values:
//     
//     "achievements" &mdash; Achievements 
//     "ribbons" &mdash; Ribbons
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Localization language. Default is "en". Valid values:
//     
//     "en" &mdash; English (by default)
//     "ru" &mdash; Русский 
//     "pl" &mdash; Polski 
//     "de" &mdash; Deutsch 
//     "fr" &mdash; Français 
//     "es" &mdash; Español 
//     "tr" &mdash; Türkçe
func (client *Client) WotxEncyclopediaAchievements(realm Realm, category []string, fields []string, language string) (*wotx.EncyclopediaAchievements, error) {
	if err := ValidateRealm(realm, []Realm{RealmWgcb}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"category": strings.Join(category, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wotx.EncyclopediaAchievements
	err := client.doGetDataRequest(realm, "/wotx/encyclopedia/achievements/", reqParam, &result)
	return result, err
}
