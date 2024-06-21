package main

import (
	"fmt"
	"log"
	"time"
	"web-crawler-assignment/crawler"
)

func main() {
	var startURL string
	fmt.Println("Enter the starting URL:")
	fmt.Scanln(&startURL)
	maxDepth := 2
	timeout := 2 * time.Second
	c := crawler.NewCrawler(startURL, maxDepth, timeout)
	if err := c.Start(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
