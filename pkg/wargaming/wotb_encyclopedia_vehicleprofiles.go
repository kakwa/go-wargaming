package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wotb"
	"strings"
)

// WotbEncyclopediaVehicleprofiles Method returns vehicle configuration characteristics.
//
// https://developers.wargaming.net/reference/all/wotb/encyclopedia/vehicleprofiles
//
// tankId:
//     Vehicle ID. Maximum limit: 25.
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
// orderBy:
//     Sorting. Valid values:
//     
//     "price_credit" &mdash; by cost in credits 
//     "-price_credit" &mdash; by cost in credits, in reverse order
func (client *Client) WotbEncyclopediaVehicleprofiles(realm Realm, tankId []int, fields []string, language string, orderBy string) (*wotb.EncyclopediaVehicleprofiles, error) {
	if err := validateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"tank_id": utils.SliceIntToString(tankId, ","),
		"fields": strings.Join(fields, ","),
		"language": language,
		"order_by": orderBy,
	}

	var result *wotb.EncyclopediaVehicleprofiles
	err := client.doGetDataRequest(realm, "/wotb/encyclopedia/vehicleprofiles/", reqParam, &result)
	return result, err
}
