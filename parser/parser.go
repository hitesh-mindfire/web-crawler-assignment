package parser

import (
	"fmt"
	"strings"
	"web-crawler-assignment/fetcher"

	"golang.org/x/net/html"
)

func Parse() []string {
	var links []string

	data, _ := fetcher.Fetch()

	htmlData, err := html.Parse(strings.NewReader(string(data)))
	if err != nil {
		return links
	}

	var extractLinks func(*html.Node)
	extractLinks = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}
			}
		}
		fmt.Println(n, "abc")
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractLinks(c)
		}
	}

	extractLinks(htmlData)

	return links
}
