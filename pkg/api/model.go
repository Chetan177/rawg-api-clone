package api

// Game Struct represent the whole data structure in mongodb
type Game struct {
	Slug             string           `bson:"slug"`
	Name             string           `bson:"name"`
	Id               int              `bson:"id"`
	Released         string           `bson:"released"`
	Tba              bool             `bson:"tba"`
	BackgroundImage  string           `bson:"background_image"`
	Rating           float64          `bson:"rating"`
	RatingTop        int              `bson:"rating_top"`
	Ratings          []Rating         `bson:"ratings"`
	RatingsCount     int              `bson:"ratings_count"`
	ReviewsTextCount int              `bson:"reviews_text_count"`
	Added            int              `bson:"added"`
	AddedByStatus    AddedByStatus    `bson:"added_by_status"`
	Metacritic       int              `bson:"metacritic"`
	Playtime         int              `bson:"playtime"`
	SuggestionsCount int              `bson:"suggestions_count"`
	Updated          string           `bson:"updated"`
	ReviewsCount     int              `bson:"reviews_count"`
	SaturatedColor   string           `bson:"saturated_color"`
	DominantColor    string           `bson:"dominant_color"`
	Platforms        []Platforms      `bson:"platforms"`
	ParentPlatforms  []ParentPlatform `bson:"parent_platforms"`
	Genres           []Genre          `bson:"genres"`
	Stores           []Store          `bson:"stores"`
	Clip             Clip             `bson:"clip"`
	Tags             []Tag            `bson:"tags"`
	ESRBRating       ESRB             `bson:"esrb_rating"`
	ShortScreenshots []Screenshot     `bson:"short_screenshots"`
}

type Rating struct {
	Id      int     `bson:"id"`
	Title   string  `bson:"title"`
	Count   int     `bson:"count"`
	Percent float64 `bson:"percent"`
}

type AddedByStatus struct {
	Yet     int `bson:"yet"`
	Owned   int `bson:"owned"`
	Beaten  int `bson:"beaten"`
	Toplay  int `bson:"toplay"`
	Dropped int `bson:"dropped"`
	Playing int `bson:"playing"`
}

type Platforms struct {
	Platform       Platform     `bson:"platform"`
	ReleasedAt     string       `bson:"released_at"`
	RequirementsEn Requirements `bson:"requirements_en"`
	RequirementsRu Requirements `bson:"requirements_ru"`
}

type Platform struct {
	Id              int    `bson:"id"`
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	Image           string `bson:"image"`
	YearEnd         int    `bson:"year_end"`
	YearStart       int    `bson:"year_start"`
	GamesCount      int    `bson:"games_count"`
	ImageBackground string `bson:"image_background"`
}

type Requirements struct {
	Minimum     string `bson:"minimum"`
	Recommended string `bson:"recommended"`
}

type ParentPlatform struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
	Slug string `bson:"slug"`
}

type Genre struct {
	Id              int    `bson:"id"`
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	GamesCount      int    `bson:"games_count"`
	ImageBackground string `bson:"image_background"`
}

type Stores struct {
	Id    int   `bson:"id"`
	Store Store `bson:"store"`
}

type Store struct {
	Id              int    `bson:"id"`
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	GamesCount      int    `bson:"games_count"`
	ImageBackground string `bson:"image_background"`
	Domain          string `bson:"domain"`
}

type Clip struct {
	Clip    string   `bson:"clip"`
	Clips   ClipsRes `bson:"clips"`
	Video   string   `bson:"video"`
	Preview string   `bson:"preview"`
}

type ClipsRes struct {
	R320 string `bson:"320"`
	R640 string `bson:"640"`
	Full string `bson:"full"`
}

type Tag struct {
	Id              int    `bson:"id"`
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	GamesCount      int    `bson:"games_count"`
	Language        string `bson:"language"`
	ImageBackground string `bson:"image_background"`
}

type ESRB struct {
	Id   int    `bson:"id"`
	Name string `bson:"name"`
	Slug string `bson:"slug"`
}

type Screenshot struct {
	Id    int    `bson:"id"`
	Image string `bson:"image"`
}
