package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgnTime"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotRatingsAccounts Method returns player ratings by specified IDs.
//
// https://developers.wargaming.net/reference/all/wot/ratings/accounts
//
// Deprecated: Attention! The method is deprecated.
//
// account_id:
//     IDs of player accounts. Maximum limit: 100.
// date:
//     Ratings calculation date. Up to 7 days before the current date; default value: yesterday. Date in UNIX timestamp or ISO 8601 format. E.g.: 1376542800 or 2013-08-15T00:00:00
// type:
//     Rating period. For valid values, check the Types of ratings method.
// battle_type:
//     Battle types. Default is "default". Valid values:
//     
//     "company" &mdash; Tank Company Battles 
//     "random" &mdash; Random Battles 
//     "team" &mdash; Team Battles 
//     "default" &mdash; any battle type (by default)
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
func (client *Client) WotRatingsAccounts(realm Realm, accountId []int, date wgnTime.UnixTime, type_ string, battleType string, fields []string, language string) (*wot.RatingsAccounts, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"account_id": utils.SliceIntToString(accountId, ","),
		"date": strconv.FormatInt(date.Unix(), 10),
		"type": type_,
		"battle_type": battleType,
		"fields": strings.Join(fields, ","),
		"language": language,
	}

	var result *wot.RatingsAccounts
	err := client.doGetDataRequest(realm, "/wot/ratings/accounts/", reqParam, &result)
	return result, err
}
