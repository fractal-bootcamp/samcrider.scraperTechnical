package scraper

import (
	"log"
	"math/rand/v2"
	"net/http"
	"samcrider/scraper/utils"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// main scraper function
func Scrape(url string, domain_name string, res *http.Response, depth int, links_per_page int) {

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// log
	log.Println("Loading HTML from:", url)

	/* clean html */
	clean_html := html_cleaner(doc)

	// set filename
	filename := utils.Parse_url(url)

	// append random number in case host name is the same
	num := rand.Int64N(1000)
	file := filename + strconv.FormatInt(num, 10)

	/* save clean html */
	utils.Create_File(file, strings.Split(clean_html, ">"))

	// base case (0, not 1, because depth will be updated on first call too)
	if depth < 1 {
		return
	}

	/* extract links */

	// get links-per-page links
	// loop through them  --> do something with the depth
	// recursively call scrape with the current link

	// EDGE CASE: the current page doesn't have at least [links_per_page] links
	current_page_links := links_per_page
	var batch_of_links []string
	doc.Find("link, a").EachWithBreak(func(i int, s *goquery.Selection) bool {
		if current_page_links < 1 {
			return false
		}
		link, exists := s.Attr("href")
		if exists {
			// check if valid url
			true_url, domain_name, res, err := utils.Test_url(link, domain_name)
			if err != nil {
				log.Println("This link: "+link, "is inaccessible")
				return true
			}

			// add valid url to current batch
			batch_of_links = append(batch_of_links, true_url)

			// scrape valid url
			Scrape(true_url, domain_name, res, depth-1, links_per_page)

			current_page_links = current_page_links - 1
		}
		return true
	})

	// log
	log.Println("Batch of links from", url+":")
	for link := range batch_of_links {
		// log
		log.Println(batch_of_links[link])
	}

}
