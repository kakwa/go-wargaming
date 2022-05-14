package wargaming

import (
	"github.com/IceflowRE/go-wargaming/pkg/wargaming/utils"
	"strings"
	"strconv"
)

type WgnWgtvVideos struct {
	// List of category IDs
	CategoryId []int `json:"category_id,omitempty"`
	// Localized description of the video
	Description string `json:"description,omitempty"`
	// Video duration in seconds. The field can return null.
	Duration int `json:"duration,omitempty"`
	// Name of the video on YouTube
	ExtTitle string `json:"ext_title,omitempty"`
	// "Important" mark
	Important bool `json:"important,omitempty"`
	// Program ID
	ProgramId int `json:"program_id,omitempty"`
	// List of project IDs
	ProjectId []int `json:"project_id,omitempty"`
	// Publication date and time
	PublishedAt UnixTime `json:"published_at,omitempty"`
	// Localized name of the video
	Title string `json:"title,omitempty"`
	// List of IDs of vehicles covered in the video
	Vehicles map[string]string `json:"vehicles,omitempty"`
	// Youtube ID
	VideoId string `json:"video_id,omitempty"`
	// Youtube link
	VideoUrl string `json:"video_url,omitempty"`
	// List of images
	Thumbnails struct {
		// Height
		Height int `json:"height,omitempty"`
		// Image link
		Url string `json:"url,omitempty"`
		// Width
		Width int `json:"width,omitempty"`
	} `json:"thumbnails,omitempty"`
}

// WgnWgtvVideos Method returns list of videos filtered by the specified parameter.
//
// https://developers.wargaming.net/reference/all/wgn/wgtv/videos
//
// category_id:
//     Content category ID. Maximum limit: 100.
// date_from:
//     Dated from the specified date
// fields:
//     Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
// important:
//     "Important" mark. Min value is 0. Maximum value: 1.
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
//     Number of returned entries (fewer can be returned, but not more than 1000). If the limit sent exceeds 1000, a limit of 100 is applied (by default).
// page_no:
//     Result page number. Default is 1. Min value is 1.
// program_id:
//     Program ID. Maximum limit: 100.
// project_id:
//     Game project ID. Maximum limit: 100.
// q:
//     Text for search by name
// vehicle_id:
//     Vehicle ID. Maximum limit: 100.
// video_id:
//     Youtube ID. Maximum limit: 100.
func (client *Client) WgnWgtvVideos(realm Realm, categoryId []int, dateFrom UnixTime, fields []string, important int, language string, limit int, pageNo int, programId []int, projectId []int, q string, vehicleId []int, videoId []string) (*WgnWgtvVideos, error) {
	if err := ValidateRealm(realm, []Realm{RealmAsia, RealmEu, RealmNa, RealmRu}); err != nil {
		return nil, err
	}

	reqParam := map[string]string{
		"category_id": utils.SliceIntToString(categoryId, ","),
		"date_from": strconv.FormatInt(dateFrom.Unix(), 10),
		"fields": strings.Join(fields, ","),
		"important": strconv.Itoa(important),
		"language": language,
		"limit": strconv.Itoa(limit),
		"page_no": strconv.Itoa(pageNo),
		"program_id": utils.SliceIntToString(programId, ","),
		"project_id": utils.SliceIntToString(projectId, ","),
		"q": q,
		"vehicle_id": utils.SliceIntToString(vehicleId, ","),
		"video_id": strings.Join(videoId, ","),
	}

	var result *WgnWgtvVideos
	err := client.doGetDataRequest(realm, "/wgn/wgtv/videos/", reqParam, &result)
	return result, err
}
