package main

import(
	"log"
	. "mgopa/hunter"
	"regexp"
)

func main() {
	log.Print("gopa is on.")

	curl := make(chan []byte)
	success := make(chan Task)
	failure := make(chan string)

	visited := make(map[string]int)

	regex := regexp.MustCompile("<a.*?href=[\"'](http.*?)[\"']")

	go Seed(curl, "http://www.baidu.com")

	go ThrottledCrawl(curl, success, failure, visited)

	for {
		site := <-success
		go GetUrls(curl, site, regex)
	}
}
