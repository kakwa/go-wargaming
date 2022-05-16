package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wot"
	"strconv"
	"strings"
)

// WotGlobalmapSeasons Method returns information about seasons.
//
// https://developers.wargaming.net/reference/all/wot/globalmap/seasons
//
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// language:
//     Language. Default is "ru". Valid values:
//     
//     "ru" &mdash; Russian (by default)
// limit:
//     Page limit. Default is 5. Min value is 1. Maximum value: 20.
// page_no:
//     Page number. Default is 1. Min value is 1.
// season_id:
//     Season ID
// status:
//     Response with seasons filtered by status. Valid values:
//     
//     "PLANNED" &mdash; Upcoming season 
//     "ACTIVE" &mdash; Current season 
//     "FINISHED" &mdash; Season is over
func (client *Client) WotGlobalmapSeasons(realm Realm, fields []string, language string, limit int, pageNo int, seasonId string, status string) (*wot.GlobalmapSeasons, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"fields": strings.Join(fields, ","),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"season_id": seasonId,
		"status": status,
	}

	var result *wot.GlobalmapSeasons
	err := client.doGetDataRequest(realm, "/wot/globalmap/seasons/", reqParam, &result)
	return result, err
}
