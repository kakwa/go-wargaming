package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/wgn"
	"strconv"
)

// WgnWargagRate Method allows to rate content. A valid access_token for the specified account is required.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/rate
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//     Parameter is required.
// contentId:
//     Publication ID
//     Parameter is required.
// rating:
//     Rate. Valid values:
//     
//     "up" - Rate content positively 
//     "down" - Rate content negatively
//     Parameter is required.
func (client *Client) WgnWargagRate(realm Realm, accessToken string, contentId int, rating string) (*wgn.WargagRate, error) {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"content_id": strconv.Itoa(contentId),
		"rating": rating,
	}

	var result *wgn.WargagRate
	err := client.doPostDataRequest(realm, "/wgn/wargag/rate/", reqParam, &result)
	return result, err
}
