package crawler

import (
	"log"
	"sync"
	"time"
	"web-crawler-assignment/fetcher"
	"web-crawler-assignment/parser"
	"web-crawler-assignment/storage"
)

type Crawler struct {
	startURL  string
	maxDepth  int
	timeout   time.Duration
	proxyUrl  string
	storage   *storage.PageStorage
	wg        sync.WaitGroup
	urlChan   chan string
	depthChan chan int
}

func NewCrawler(startURL string, maxDepth int, timeout time.Duration, proxyUrl string, jsonOutput bool) *Crawler {
	return &Crawler{
		startURL:  startURL,
		maxDepth:  maxDepth,
		timeout:   timeout,
		proxyUrl:  proxyUrl,
		storage:   storage.NewPageStorage(jsonOutput),
		urlChan:   make(chan string),
		depthChan: make(chan int),
	}
}

func (c *Crawler) Start() error {
	log.Println("Start crawler", c)

	c.wg.Add(1)
	go c.crawl(c.startURL, 0)

	go func() {
		for url := range c.urlChan {
			depth := <-c.depthChan
			if depth <= c.maxDepth && !c.storage.HasVisited(url) {
				c.wg.Add(1)
				go c.crawl(url, depth)
			}
		}
	}()

	c.wg.Wait()
	log.Println("Finished crawler", c)
	if c.storage.IsJSONOutput() {
		if err := c.storage.WriteJSONToFile("crawler_results.json"); err != nil {
			log.Println("Error writing JSON to file:", err)
		}
	}

	return nil
}

func (c *Crawler) crawl(url string, depth int) {
	defer c.wg.Done()

	if depth > c.maxDepth || c.storage.HasVisited(url) {
		return
	}

	c.storage.MarkVisited(url)

	data, err := fetcher.Fetch(url, c.timeout, c.proxyUrl)
	if err != nil {
		log.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}

	c.storage.StoreContent(url, data)
	links := parser.Parse(data)
	for _, link := range links {
		if !c.storage.HasVisited(link) {
			c.urlChan <- link
			c.depthChan <- depth + 1
		}
	}
}
