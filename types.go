package gochat

type Response struct {
	Name   string `json:"name"`
	Text   string `json:"text"`
	Thread Space  `json:"thread"`
	Space  Space  `json:"space"`
}

type Space struct {
	Name string `json:"name"`
}
