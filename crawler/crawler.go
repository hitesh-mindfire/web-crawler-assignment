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
	startURL         string
	maxDepth         int
	timeout          time.Duration
	proxyUrl         string
	maxSize          int
	disableRedirects bool
	showSource       bool
	insecure         bool
	uniqueUrls       bool
	storage          *storage.PageStorage
	wg               sync.WaitGroup
	urlChan          chan string
	depthChan        chan int
}

func NewCrawler(startURL string, maxDepth int, timeout time.Duration, proxyUrl string, jsonOutput bool, maxSize int, disableRedirects bool, showSource bool, insecure bool, uniqueUrls bool) *Crawler {
	return &Crawler{
		startURL:         startURL,
		maxDepth:         maxDepth,
		timeout:          timeout,
		proxyUrl:         proxyUrl,
		maxSize:          maxSize,
		disableRedirects: disableRedirects,
		showSource:       showSource,
		insecure:         insecure,
		uniqueUrls:       uniqueUrls,
		storage:          storage.NewPageStorage(jsonOutput, maxSize),
		urlChan:          make(chan string),
		depthChan:        make(chan int),
	}
}

func (c *Crawler) Start() error {
	log.Println("Start crawler", c)

	c.wg.Add(1)
	go c.crawl(c.startURL, 0)
	c.storage.StoreSource(c.startURL, "href")

	go func() {
		for url := range c.urlChan {
			depth := <-c.depthChan
			if depth <= c.maxDepth && (!c.uniqueUrls || !c.storage.HasVisited(url)) {
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

	if depth > c.maxDepth {
		return
	}

	if c.uniqueUrls {
		c.storage.MarkVisited(url)
	}

	data, size, err := fetcher.Fetch(url, c.timeout, c.proxyUrl, c.disableRedirects, c.insecure)
	if err != nil {
		log.Printf("Error fetching URL %s: %v\n", url, err)
		return
	}

	if c.maxSize > 0 && size > c.maxSize*1024 {
		log.Printf("Skipping URL %s due to size limit (%d bytes > %d bytes)\n", url, size, c.maxSize*1024)
		return
	}

	c.storage.StoreContent(url, data, c.showSource)

	links := parser.Parse(data)
	for link, source := range links {
		if !c.uniqueUrls || !c.storage.HasVisited(link) {
			c.urlChan <- link
			c.depthChan <- depth + 1
			c.storage.StoreSource(link, source)
		}
	}
}
