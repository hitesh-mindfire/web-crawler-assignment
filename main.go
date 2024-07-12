package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
	"web-crawler-assignment/crawler"
)

func main() {
	url := flag.String("url", "", "The starting URL")
	depth := flag.Int("depth", 2, "The maximum depth to crawl")
	timeout := flag.Int("timeout", 2, "The timeout for each request in seconds")
	proxy := flag.String("proxy", "", "Proxy URL. E.g. -proxy http://127.0.0.1:8080")
	jsonOutput := flag.Bool("json", false, "Output as JSON")
	showHelp := flag.Bool("h", false, "")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s web crawler:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	fmt.Println(showHelp)
	if *showHelp {
		flag.Usage()
		return
	}

	if *url == "" {
		log.Fatal("You must specify a starting URL with the -u flag")
	}

	timeoutDuration := time.Duration(*timeout) * time.Second

	c := crawler.NewCrawler(*url, *depth, timeoutDuration, *proxy, *jsonOutput)

	if err := c.Start(); err != nil {
		log.Fatalf("error: %v", err)
	}
}
