package main

import (
	"log"
	"time"
	"web-crawler-assignment/crawler"
)

func main() {
	startURL := "https://go.dev/doc/"
	maxDepth := 2
	timeout := 2 * time.Second
	c := crawler.NewCrawler(startURL, maxDepth, timeout)
	if err := c.Start(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
