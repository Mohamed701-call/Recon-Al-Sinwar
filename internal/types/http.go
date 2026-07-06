package types

type HTTPResult struct {
	URL        string   `json:"url"`
	Input      string   `json:"input"`
	Host       string   `json:"host"`
	IP         string   `json:"ip"`
	Port       string   `json:"port"`
	WebServer  string   `json:"webserver"`
	Title      string   `json:"title"`
	StatusCode int      `json:"status_code"`
	Tech       []string `json:"tech"`
	CDN        string   `json:"cdn"`
}
