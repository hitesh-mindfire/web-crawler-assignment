package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func Parse(body []byte) map[string]string {
	links := make(map[string]string)
	htmlData, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return links
	}

	links = extractLinks(htmlData, links)
	return links
}

func extractLinks(n *html.Node, links map[string]string) map[string]string {
	if n.Type == html.ElementNode {
		if n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links[attr.Val] = "href"
				}
			}
		}
		if n.Data == "script" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links[attr.Val] = "script"
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = extractLinks(c, links)
	}
	return links
}
