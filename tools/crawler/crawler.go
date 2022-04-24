package crawler

import (
	"net/http"

	"github.com/jackdanger/collectlinks"
)

func Crawl(url string) []string {

	resp, _ := http.Get(url)
	return collectlinks.All(resp.Body)
}
