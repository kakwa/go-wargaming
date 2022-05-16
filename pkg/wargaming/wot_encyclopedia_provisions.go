package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotEncyclopediaProvisions Method returns a list of available equipment and consumables.
//
// https://developers.wargaming.net/reference/all/wot/encyclopedia/provisions
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
// limit:
//     Number of returned entries (fewer can be returned, but not more than 100). If the limit sent exceeds 100, a limit of 100 is applied (by default).
// page_no:
//     Result page number
// provision_id:
//     Equipment or consumables ID. Maximum limit: 100.
// type:
//     Type. Default is "equipment, optionalDevice". Maximum limit: 100. Valid values:
//     
//     "equipment" &mdash; Consumables 
//     "optionalDevice" &mdash; Equipment
func (client *Client) WotEncyclopediaProvisions(realm Realm, fields []string, language string, limit int, pageNo int, provisionId []int, type_ []string) (*wot.EncyclopediaProvisions, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"provision_id": utils.SliceIntToString(provisionId, ","),
		"type": strings.Join(type_, ","),
	}

	var result *wot.EncyclopediaProvisions
	err := client.doGetDataRequest(realm, "/wot/encyclopedia/provisions/", reqParam, &result)
	return result, err
}
