package utils

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	parser "net/url"
	"os"
	"strconv"
	"strings"
)

func Create_File(name string, file_content []string) {
	// create file
	file, err := os.Create("./scraper/clean_html/" + name + ".html")
	if err != nil {
		fmt.Println(err)
		return
	}

	// makes sure the file closes when function finishes execution
	defer closeFile(file)

	// loop through data and write lines
	for _, v := range file_content {
		_, err := fmt.Fprintln(file, v)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	/* Log file size */
	fi, err := os.Stat("./scraper/clean_html/" + name + ".html")
	if err != nil {
		log.Fatal(err)
	}
	// get the size in kilobytes
	size := fi.Size() / 1024

	fmt.Println("Size of " + name + ": " + strconv.FormatInt(size, 10))
}

func Arg_parser(args []string) (string, int, int) {
	// store url in variable
	url := args[0]
	if url == "" {
		log.Fatal("No url provided")
	}

	// init depth
	var depth int
	var err error
	// search args for --depth flag
	for i := range args {
		if args[i] == "--depth" {
			// Found, set depth
			depth, err = strconv.Atoi(args[i+1])
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	// init links per page
	var links_per_page int
	// search args for --per-page flag
	for i := range args {
		if args[i] == "--per-page" {
			// Found, set links per page
			links_per_page, err = strconv.Atoi(args[i+1])
			if err != nil {
				log.Fatal(err)
			}
			break
		}
	}

	return url, depth, links_per_page
}

func Parse_domain_name(link string) string {
	var domain_name string

	domain_name_array := strings.Split(link, "")
	// counter to count '/'
	counter := 0
	for char := range domain_name_array {
		if domain_name_array[char] == "/" {
			if counter == 2 {
				domain_name = strings.Join(domain_name_array[:char], "")
				log.Println(domain_name)
				break
			}
			counter = counter + 1
		}
	}
	return domain_name
}

func Test_url(link string, domain_name string) (string, string, *http.Response, error) {

	// check if valid url
	ok_url, err := parser.ParseRequestURI(link)
	if err != nil {
		return "", "", nil, err
	}

	// get valid url as a string
	valid_url := ok_url.String()

	// check if '/' is first character
	urlArray := strings.Split(valid_url, "")
	// if yes, append domain name to front
	if urlArray[0] == "/" {
		valid_url = domain_name + valid_url
	} else {
		domain_name = Parse_domain_name(valid_url)
	}

	// test link
	res, err := http.Get(valid_url)
	if err != nil {
		return "", "", nil, err
	}

	// check status code
	if res.StatusCode != 200 {
		return "", "", nil, errors.New("status code error: " + strconv.Itoa(res.StatusCode) + res.Status)
	}

	// return the response of the get request
	return valid_url, domain_name, res, nil
}

func Parse_url(u string) string {
	url, err := parser.Parse(u)
	if err != nil {
		log.Println("url couldn't be parsed")
		return "filename"
	}
	var urlstr string = url.String()
	var hostname string
	var temp []string

	if strings.HasPrefix(urlstr, "https") {
		hostname = strings.TrimPrefix(urlstr, "https://")
	} else if strings.HasPrefix(urlstr, "http") {
		hostname = strings.TrimPrefix(urlstr, "http://")
	} else {
		hostname = urlstr
	}

	if strings.HasPrefix(hostname, "www") {
		hostname = strings.TrimPrefix(hostname, "www.")
	}
	if strings.Contains(hostname, ".") {
		temp = strings.Split(hostname, ".")
		hostname = temp[0]
	}
	if strings.Contains(hostname, "/") {
		temp = strings.Split(hostname, "/")
		hostname = temp[0]
	}

	return hostname
}
