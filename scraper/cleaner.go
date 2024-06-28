package scraper

import (
	"log"

	"github.com/PuerkitoBio/goquery"
)

// html cleaner function
func html_cleaner(doc *goquery.Document) string {
	// remove script tags
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// remove nav tags
	doc.Find("nav").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// remove elements with class .vector-header
	doc.Find(".vector-header").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// remove elements with ID #p-lang-btn
	doc.Find("#p-lang-btn").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// remove elements with the class .infobox
	doc.Find(".infobox").Each(func(i int, s *goquery.Selection) {
		s.Remove()
	})

	// remove stylesheet links
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		stylesheet, exists := s.Attr("rel")
		if exists {
			if stylesheet == "stylesheet" {
				s.Remove()
			}
		}
	})

	// get the doc in html and return it
	html_content, err := doc.Html()
	if err != nil {
		log.Fatal(err)
	}

	return html_content
}
