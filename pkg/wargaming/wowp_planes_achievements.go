package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strconv"
	"strings"
)

// WowpPlanesAchievements Method returns achievements on player's aircraft.
//
// https://developers.wargaming.net/reference/all/wowp/planes/achievements
//
// accountId:
//     Player account ID
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
// planeId:
//     Aircraft ID. Maximum limit: 100.
func (client *Client) WowpPlanesAchievements(realm Realm, accountId int, fields []string, language string, planeId []int) (*wowp.PlanesAchievements, error) {
	if err := validateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": strconv.Itoa(accountId),
		"fields": strings.Join(fields, ","),
		"language": language,
		"plane_id": utils.SliceIntToString(planeId, ","),
	}

	var result *wowp.PlanesAchievements
	err := client.doGetDataRequest(realm, "/wowp/planes/achievements/", reqParam, &result)
	return result, err
}
