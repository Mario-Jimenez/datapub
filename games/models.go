package games

type (
	Game struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		BasePrice   float32 `json:"base_price"`
		ScoreURL    string  `json:"score_url"`
		DurationURL string  `json:"duration_url"`
		PricesURLs  []price `json:"prices_urls"`
	}

	price struct {
		From string `json:"from"`
		URL  string `json:"url"`
	}
)
