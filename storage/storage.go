package storage

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

type PageStorage struct {
	visitedUrls map[string]bool
	pageContent map[string][]byte
	urlSource   map[string]string
	jsonOutput  bool
	mutex       sync.Mutex
	urls        []string
	maxSize     int
}

func NewPageStorage(jsonOutput bool, maxSize int) *PageStorage {
	return &PageStorage{
		visitedUrls: make(map[string]bool),
		pageContent: make(map[string][]byte),
		urlSource:   make(map[string]string),
		jsonOutput:  jsonOutput,
		maxSize:     maxSize,
		urls:        []string{},
	}
}

func (ps *PageStorage) MarkVisited(url string) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.visitedUrls[url] = true
}

func (ps *PageStorage) HasVisited(url string) bool {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	return ps.visitedUrls[url]
}

func (ps *PageStorage) StoreSource(url, source string) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.urlSource[url] = source
}

func (ps *PageStorage) GetContent(url string) ([]byte, bool) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	content, exists := ps.pageContent[url]
	return content, exists
}

func (ps *PageStorage) IsJSONOutput() bool {
	return ps.jsonOutput
}

func (ps *PageStorage) StoreContent(url string, content []byte, showSource bool) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.pageContent[url] = content

	if ps.jsonOutput {
		ps.urls = append(ps.urls, url)
	} else {
		source := ps.urlSource[url]
		err := storeUrlToFile("crawler_results.txt", url, showSource, source)
		if err != nil {
			log.Println(err)
		}
	}
}

func storeUrlToFile(filename, url string, showSource bool, source string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if showSource && source != "" {
		_, err = file.WriteString("[" + source + "] " + url + "\n")
	} else {
		_, err = file.WriteString(url + "\n")
	}
	return err
}

func (ps *PageStorage) WriteJSONToFile(filename string) error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	var urlsWithSources []string
	for _, url := range ps.urls {
		if source, found := ps.urlSource[url]; found {
			urlsWithSources = append(urlsWithSources, "["+source+"] "+url)
		} else {
			urlsWithSources = append(urlsWithSources, url)
		}
	}

	data := map[string]interface{}{
		"urls": urlsWithSources,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}
