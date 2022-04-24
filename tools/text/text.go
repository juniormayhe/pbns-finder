package text

import (
	"regexp"
	"strings"
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
