package wargaming

import (
	"strconv"
)

// WgnWargagDeletecomment Method allows to delete comments. A valid access_token for the specified account is required.
//
// https://developers.wargaming.net/reference/all/wgn/wargag/deletecomment
//
// accessToken:
//     Access token for the private data of a user's account; can be received via the authorization method; valid within a stated time period
//     Parameter is required.
// commentId:
//     Comment ID
//     Parameter is required.
func (client *Client) WgnWargagDeletecomment(realm Realm, accessToken string, commentId int) error {
	if err := validateRealm(realm, []Realm{RealmRu}); err != nil {
		return nil
	}

	reqParam := map[string]string{
		"access_token": accessToken,
		"comment_id": strconv.Itoa(commentId),
	}

	err := client.doPostRequest(realm, "/wgn/wargag/deletecomment/", reqParam)
	return err
}
