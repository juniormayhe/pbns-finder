package main

import (
	"fmt"
	"regexp"
	"juniormayhe.com/pbns-finder/tools"
)

func main() {
	text.PrintHello()
	r := regexp.MustCompile(`[\w]+\-[\w]+\-[\w]*`)
	matches := r.FindAllString("um-dois, commerce-infra-teste, -alguma\n  coisa, commerce-catalog-app, - alguma\n\tcoisa ", -1)
	fmt.Println(len(matches))
	fmt.Println(matches)
	for _, v := range matches {
		fmt.Println("valor=", v)
	}

}
