// Auto generated file!

package wot

import (
	"github.com/IceflowRE/go-wargaming/v4/wargaming/wgnTime"
)

type AuthLoginOptions struct {
	// Layout for mobile applications. Valid values:
	//
	// "page" - Page
	// "popup" - Popup window
	Display *string `json:"display,omitempty"`
	// Access_token expiration time in UNIX. Delta can also be specified in seconds.
	// Expiration time and delta must not exceed two weeks from the current time.
	ExpiresAt *wgnTime.UnixTime `json:"expires_at,omitempty"`
	// If parameter nofollow=1 is passed in, the user is not redirected. URL is returned in response. Default is 0. Min value is 0. Maximum value: 1.
	Nofollow *int `json:"nofollow,omitempty"`
	// URL where user is redirected after authentication.
	// By default: api.worldoftanks.eu/wot	//blank/
	RedirectUri *string `json:"redirect_uri,omitempty"`
}

type AuthLogin struct {
	// URL where user is redirected for authentication.
	// This URL is returned only if parameter nofollow=1 is passed in.
	Location *string `json:"location,omitempty"`
}
