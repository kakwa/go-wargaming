package wot


// GlobalmapSeasonsOptions options.
type GlobalmapSeasonsOptions struct {
	// Response field. The fields are separated with commas. Embedded fields are separated with dots. To exclude a field, use “-” in front of its name. In case the parameter is not defined, the method returns all fields. Maximum limit: 100.
	Fields []string
	// Language. Default is "ru". Valid values:
	// 
	// "ru" - Russian (by default)
	Language *string
	// Page limit. Default is 5. Min value is 1. Maximum value: 20.
	Limit *int
	// Page number. Default is 1. Min value is 1.
	PageNo *int
	// Season ID
	SeasonId *string
	// Response with seasons filtered by status. Valid values:
	// 
	// "PLANNED" - Upcoming season 
	// "ACTIVE" - Current season 
	// "FINISHED" - Season is over
	Status *string
}

type GlobalmapSeasons struct {
	// Finishing time
	End *string `json:"end,omitempty"`
	// Fronts. Only season fronts are presented in response.
	Fronts []*struct {
		// Front ID
		FrontId *string `json:"front_id,omitempty"`
		// Front name
		FrontName *string `json:"front_name,omitempty"`
		// Link to Front
		Url *string `json:"url,omitempty"`
	} `json:"fronts,omitempty"`
	// Season ID
	SeasonId *string `json:"season_id,omitempty"`
	// Season name
	SeasonName *string `json:"season_name,omitempty"`
	// Start time
	Start *string `json:"start,omitempty"`
	// Status
	Status *string `json:"status,omitempty"`
}
