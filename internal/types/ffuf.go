package types

type FuzzTarget struct {
	URL   string
	Score int
}

type FFUFResult struct {
	Results []FFUFEntry `json:"results"`
}

type FFUFEntry struct {
	URL         string `json:"url"`
	Status      int    `json:"status"`
	Length      int    `json:"length"`
	Words       int    `json:"words"`
	Lines       int    `json:"lines"`
	ContentType string `json:"content-type"`
}
