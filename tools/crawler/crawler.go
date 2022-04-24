package crawler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/jackdanger/collectlinks"
)

func GetBodyContent(sourceUrl string) (string, error) {
	resp, err := http.Get(sourceUrl)

	if err != nil || resp.StatusCode != http.StatusOK {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func GetUrlsWithExtension(sourceUrl, extension string) []string {

	resp, err := http.Get(sourceUrl)
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.StatusCode)
		return []string{}
	}

	urls := collectlinks.All(resp.Body)
	var filteredUrls []string

	for _, item := range urls {
		if strings.HasSuffix(item, extension) {

			if !strings.HasPrefix(item, "http") {
				u, _ := url.Parse(sourceUrl)
				filteredUrls = append(filteredUrls, u.Scheme+"://"+u.Host+""+item)
			} else {
				filteredUrls = append(filteredUrls, item)
			}

		}
	}

	return filteredUrls

}

func ReplaceGithubBlobWithRawContent(url string) string {
	s := strings.Replace(url, "/blob/main/", "/main/", -1)
	s = strings.Replace(s, "github.com", "raw.githubusercontent.com", -1)
	return s
}
