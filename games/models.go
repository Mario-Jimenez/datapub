package games

type Game struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	ScoreURL    string   `json:"score_url"`
	DurationURL string   `json:"duration_url"`
	PricesURLs  []string `json:"prices_urls"`
}
