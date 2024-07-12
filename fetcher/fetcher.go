package fetcher

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func Fetch(targetUrl string, timeout time.Duration, proxyUrl string) ([]byte, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	if proxyUrl != "" {
		proxy, err := url.Parse(proxyUrl)
		if err != nil {
			return nil, err
		}
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	resp, err := client.Get(targetUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
