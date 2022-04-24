package main

import (
	"fmt"

	crawler "juniormayhe.com/pbns-finder/tools/crawler"
	text "juniormayhe.com/pbns-finder/tools/text"
)

func main() {
	url := "https://github.com/juniormayhe/pbns-finder/tree/main/yamls"
	urls := crawler.GetUrlsWithExtension(url, ".yaml")

	var body string
	var rawContentUrl string
	var err error
	for _, url := range urls {

		rawContentUrl = crawler.ReplaceGithubBlobWithRawContent(url)
		//log.Printf("fetching %s\n", rawContentUrl)

		body, err = crawler.GetBodyContent(rawContentUrl)
		if err != nil {
			fmt.Println(err)
			return
		}

		matches := text.GetPBNs(body)

		//log.Printf("pbns: %d\n", len(matches))

		fmt.Printf("file: %s\n", text.GetFilename(url))
		for _, v := range matches {
			fmt.Println("  valor: ", v)
		}
	}

}
