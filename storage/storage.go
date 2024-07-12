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
	jsonOutput  bool
	mutex       sync.Mutex
	urls        []string
}

func NewPageStorage(jsonOutput bool) *PageStorage {
	return &PageStorage{
		visitedUrls: make(map[string]bool),
		pageContent: make(map[string][]byte),
		jsonOutput:  jsonOutput,
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

func (ps *PageStorage) StoreContent(url string, content []byte) {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()
	ps.pageContent[url] = content

	if ps.jsonOutput {
		ps.urls = append(ps.urls, url)
	} else {
		err := storeUrlToFile("crawler_results.txt", url)
		if err != nil {
			log.Println(err)
		}
	}
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

func (ps *PageStorage) WriteJSONToFile(filename string) error {
	ps.mutex.Lock()
	defer ps.mutex.Unlock()

	data := map[string]interface{}{
		"urls": ps.urls,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func storeUrlToFile(filename, url string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("URL: " + url + "\n")
	return err
}
