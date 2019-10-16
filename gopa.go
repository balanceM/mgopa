package main

import(
	"flag"
	"log"
	. "mgopa/hunter"
	"regexp"
)

var seed_url = flag.String("seed", "", "Seed URL")

func main() {
	log.Print("[gopa] is on.")

	flag.Parse()
	if *seed_url == "" {
		log.Fatal("no seed was given")
	}

	curl := make(chan []byte)
	success := make(chan Task)
	failure := make(chan string)

	visited := make(map[string]int)

	regex := regexp.MustCompile("<a.*?href=[\"'](http.*?)[\"']")

	go Seed(curl, *seed_url)

	go ThrottledCrawl(curl, success, failure, visited)

	for {
		site := <-success
		go GetUrls(curl, site, regex)
	}
}
