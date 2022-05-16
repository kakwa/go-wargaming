package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wowp"
	"strings"
)

// WowpEncyclopediaPlanemodules Method returns information from Encyclopedia about modules available for specified aircrafts.
//
// https://developers.wargaming.net/reference/all/wowp/encyclopedia/planemodules
//
// plane_id:
//     Aircraft ID. Maximum limit: 1000.
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
// type:
//     Configuration. Default is "engine, bomb, rocket, turret, turretfront, turretupper, turretlower, gun, construction". Maximum limit: 100. Valid values:
//     
//     "engine" &mdash; engine 
//     "bomb" &mdash; bomb 
//     "rocket" &mdash; rocket 
//     "turret" &mdash; rear gun 
//     "turretfront" &mdash; Передняя турель 
//     "turretupper" &mdash; Верхняя турель 
//     "turretlower" &mdash; Нижняя турель 
//     "gun" &mdash; autocannon 
//     "construction" &mdash; airframe
func (client *Client) WowpEncyclopediaPlanemodules(realm Realm, planeId []int, fields []string, language string, type_ []string) (*wowp.EncyclopediaPlanemodules, error) {
	if err := ValidateRealm(realm, []Realm{RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"plane_id": utils.SliceIntToString(planeId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"type": strings.Join(type_, ","),
	}

	var result *wowp.EncyclopediaPlanemodules
	err := client.doGetDataRequest(realm, "/wowp/encyclopedia/planemodules/", reqParam, &result)
	return result, err
}
