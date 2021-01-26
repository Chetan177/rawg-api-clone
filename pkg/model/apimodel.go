package model

type GameRequest struct {
	Page              int    `json:"page"`
	PageSize          int    `json:"page_size"`
	Search            string `json:"search"`
	SearchPrecise     bool   `json:"search_precise"`
	SearchExact       bool   `json:"search_exact"`
	ParentPlatform    string `json:"parent_platform"`
	Platforms         string `json:"platforms"`
	Stores            string `json:"stores"`
	Developers        string `json:"developers"`
	Publishers        string `json:"publishers"`
	Genres            string `json:"genres"`
	Creators          string `json:"creators"`
	Dates             string `json:"dates"`
	Updated           string `json:"updated"`
	PlatformsCount    int    `json:"platforms_count"`
	Metacritic        string `json:"metacritic"`
	ExcludeCollection int    `json:"exclude_collection"`
	ExcludeAddition   bool   `json:"exclude_addition"`
	ExcludeParents    bool   `json:"exclude_parents"`
	ExcludeGameSeries bool   `json:"exclude_game_series"`
	Ordering          string `json:"ordering"`
}
