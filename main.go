package main

import (
	"github.com/gocolly/colly"
	"log"
	"os"
	"bufio"
	"fmt"
)

func main() {

	//domain := colly.AllowedDomains("oxylabs.io")

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0"),

	)
	c.OnHTML("body", func(e *colly.HTMLElement) {

		links :=  e.ChildAttrs("a", "href")

		writeFile(links)
	})

	err := c.Visit("https://www.ebay-kleinanzeigen.de/")

	if err != nil {
		log.Fatal(err)
	}
}

func writeFile(links []string) {

	file, error := os.Create("data.txt")

	if error != nil {
		log.Fatal(error)
	}

	defer file.Close()

	w := bufio.NewWriter(file)

	for _, line := range links {
		fmt.Fprintln(w, line)
	}
}