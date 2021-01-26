package model

type GameRequest struct {
	Page              int    `query:"page"`
	PageSize          int    `query:"page_size"`
	Search            string `query:"search"`
	SearchPrecise     bool   `query:"search_precise"`
	SearchExact       bool   `query:"search_exact"`
	ParentPlatform    string `query:"parent_platform"`
	Platforms         string `query:"platforms"`
	Stores            string `query:"stores"`
	Developers        string `query:"developers"`
	Publishers        string `query:"publishers"`
	Genres            string `query:"genres"`
	Creators          string `query:"creators"`
	Dates             string `query:"dates"`
	Updated           string `query:"updated"`
	PlatformsCount    int    `query:"platforms_count"`
	Metacritic        string `query:"metacritic"`
	ExcludeCollection int    `query:"exclude_collection"`
	ExcludeAddition   bool   `query:"exclude_addition"`
	ExcludeParents    bool   `query:"exclude_parents"`
	ExcludeGameSeries bool   `query:"exclude_game_series"`
	Ordering          string `query:"ordering"`
}

type Response struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
