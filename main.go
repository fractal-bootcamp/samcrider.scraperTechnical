package main

import (
	"fmt"
	"log"
	"os"
	"samcrider/scraper/scraper"
	"samcrider/scraper/utils"
)

func main() {

	// get CLI args
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("You must specify CLI arguments")
	}

	// parse CLI args
	url, depth, links_per_page := utils.Arg_parser(args)

	// define initial domain name
	domain_name := utils.Parse_domain_name(url)

	// test the url
	true_url, domain_name, res, err := utils.Test_url(url, domain_name)
	if err != nil {
		log.Fatal("This link: " + url + "is inaccessible")
	}

	// scrape the url
	scraper.Scrape(true_url, domain_name, res, depth, links_per_page)

	fmt.Println("Job completed")
}
