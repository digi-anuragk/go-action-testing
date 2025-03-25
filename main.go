package main

import (
	"fmt"
	"os"
)

func main() {
	flag := os.Args[1] == "true"

	if flag == true {
		fmt.Println("Good to go...")
	} else {
		panic(`failed to do request: all hosts have been contacted unsuccessfully, it can either be a server or a network error or wrong appID/key credentials were used. cannot perform request:
		error=Post "https://***.algolia.net/1/indexes/dev_DocsCrawler.tmp/batch": context deadline exceeded
		method=POST
		url=https://***.algolia.net/1/indexes/dev_DocsCrawler.tmp/batch
	cannot perform request:
		error=Post "https://***-3.algolianet.com/1/indexes/dev_DocsCrawler.tmp/batch": http: ContentLength=1634260 with Body length 0
		method=POST
		url=https://***-3.algolianet.com/1/indexes/dev_DocsCrawler.tmp/batch
	cannot perform request:
		error=Post "https://***-1.algolianet.com/1/indexes/dev_DocsCrawler.tmp/batch": http: ContentLength=1634260 with Body length 0
		method=POST
		url=https://***-1.algolianet.com/1/indexes/dev_DocsCrawler.tmp/batch
	cannot perform request:
		error=Post "https://***-2.algolianet.com/1/indexes/dev_DocsCrawler.tmp/batch": http: ContentLength=1634260 with Body length 0
		method=POST
		url=https://***-2.algolianet.com/1/indexes/dev_DocsCrawler.tmp/batch`)
	}
}
