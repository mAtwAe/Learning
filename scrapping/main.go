package main

import "github.com/gocolly/colly"

func scrape(url string, ch chan<- string) {
	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		ch <- e.Text
	})
	c.Visit(url)
}
