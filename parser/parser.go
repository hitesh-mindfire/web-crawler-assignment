package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func Parse(body []byte) []string {
	var links []string
	htmlData, err := html.Parse(strings.NewReader(string(body)))
	if err != nil {
		return links
	}

	return extractLinks(htmlData, links)
}

func extractLinks(n *html.Node, links []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = extractLinks(c, links)
	}
	return links
}
