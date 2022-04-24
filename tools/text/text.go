package text

import (
	"log"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

func GetPBNs(content string) []string {
	r := regexp.MustCompile(`[\w]+\-[\w]+\-[\w]*`)
	matches := r.FindAllString(content, -1)
	return matches
}

func GetFilename(url string) string {
	s := strings.Split(url, "/")
	return s[len(s)-1]
}

func GetYamlContent(s string) map[interface{}]interface{} {
	data := make(map[interface{}]interface{})
	err2 := yaml.Unmarshal([]byte(s), &data)

	if err2 != nil {

		log.Fatal(err2)
	}

	return data
}
