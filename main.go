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

	pbnsByFileName := make(map[string][]string)
	for _, url := range urls {

		rawContentUrl = crawler.ReplaceGithubBlobWithRawContent(url)

		body, err = crawler.GetBodyContent(rawContentUrl)
		if err != nil {
			fmt.Println(err)
			return
		}

		yamlMap := text.GetYamlContent(body)
		var pbns []string

		for k, v := range yamlMap {
			pbnsOnKeys := text.GetPBNs(fmt.Sprintf("%s", k))
			pbnsOnValues := text.GetPBNs(fmt.Sprintf("%s", v))

			if len(pbnsOnKeys) > 0 {
				pbns = append(pbns, pbnsOnKeys...)
			}

			if len(pbnsOnValues) > 0 {
				pbns = append(pbns, pbnsOnValues...)
			}
		}

		pbnsByFileName[text.GetFilename(url)] = pbns
	}

	for k, v := range pbnsByFileName {
		fmt.Printf("File %s contains %s\n", k, v)
	}

}
