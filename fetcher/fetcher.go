package fetcher

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func Fetch(targetUrl string, timeout time.Duration, proxyUrl string) ([]byte, int, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, 0, err
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	resp, err := client.Get(targetUrl)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, len(body), nil
}
