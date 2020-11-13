package internal

import "net/http"

// Event is used for anonymous usage statistics.
func Event(name string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://www.google-analytics.com/g/collect?v=2&tid=G-7QZCNCZ353&en=cli_"+name, nil)
	req.Header.Set("User-Agent", "PTerm/CLI")
	_, _ = client.Do(req)
}
