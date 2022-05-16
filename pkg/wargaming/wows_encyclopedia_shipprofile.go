package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wows"
	"strconv"
	"strings"
)

// WowsEncyclopediaShipprofile Method returns parameters of ships in all existing configurations.
//
// https://developers.wargaming.net/reference/all/wows/encyclopedia/shipprofile
//
// artillery_id:
//     Main Battery ID. If the module is not indicated, module of basic configuration is used.
// dive_bomber_id:
//     Dive bombers' ID. If the module is not indicated, module of basic configuration is used.
// engine_id:
//     Engine ID. If the module is not indicated, module of basic configuration is used.
// fighter_id:
//     Fighters' ID. If the module is not indicated, module of basic configuration is used.
// ship_id:
//     Ship ID
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// fire_control_id:
//     ID of Gun Fire Control System. If the module is not indicated, module of basic configuration is used.
// flight_control_id:
//     ID of Flight Control System. If the module is not indicated, module of basic configuration is used.
// hull_id:
//     Hull ID. If the module is not indicated, module of basic configuration is used.
// language:
//     Localization language. Default is "ru". Valid values:
//     
//     "cs" &mdash; Čeština 
//     "de" &mdash; Deutsch 
//     "en" &mdash; English 
//     "es" &mdash; Español 
//     "fr" &mdash; Français 
//     "ja" &mdash; 日本語 
//     "pl" &mdash; Polski 
//     "ru" &mdash; Русский (by default)
//     "th" &mdash; ไทย 
//     "zh-tw" &mdash; 繁體中文 
//     "tr" &mdash; Türkçe 
//     "zh-cn" &mdash; 中文 
//     "pt-br" &mdash; Português do Brasil 
//     "es-mx" &mdash; Español (México)
// torpedo_bomber_id:
//     Torpedo bombers' ID. If the module is not indicated, module of basic configuration is used.
// torpedoes_id:
//     Torpedo tubes' ID. If the module is not indicated, module of basic configuration is used.
func (client *Client) WowsEncyclopediaShipprofile(realm Realm, artilleryId int, diveBomberId int, engineId int, fighterId int, shipId int, fields []string, fireControlId int, flightControlId int, hullId int, language string, torpedoBomberId int, torpedoesId int) (*wows.EncyclopediaShipprofile, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"artillery_id": strconv.Itoa(artilleryId),
		"dive_bomber_id": strconv.Itoa(diveBomberId),
		"engine_id": strconv.Itoa(engineId),
		"fighter_id": strconv.Itoa(fighterId),
		"ship_id": strconv.Itoa(shipId),
		"fields": strings.Join(fields, ","),
		"fire_control_id": strconv.Itoa(fireControlId),
		"flight_control_id": strconv.Itoa(flightControlId),
		"hull_id": strconv.Itoa(hullId),
		"language": language,
		"torpedo_bomber_id": strconv.Itoa(torpedoBomberId),
		"torpedoes_id": strconv.Itoa(torpedoesId),
	}

	var result *wows.EncyclopediaShipprofile
	err := client.doGetDataRequest(realm, "/wows/encyclopedia/shipprofile/", reqParam, &result)
	return result, err
}
