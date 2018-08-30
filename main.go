package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	var newsHeadings []string

	// Find and visit all links
	c.OnHTML("a[href]>span", func(e *colly.HTMLElement) {
		if len(strings.Fields(e.Text)) > 1 {
			fmt.Println(e.Text)
			newsHeadings = append(newsHeadings, e.Text)
		}
	})

	c.Visit("https://news.google.com/")
	fmt.Println(newsHeadings)
}
