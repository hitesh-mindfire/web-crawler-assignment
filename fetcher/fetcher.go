package fetcher

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func Fetch(targetUrl string, timeout time.Duration, proxyUrl string, disableRedirects bool, insecure bool) ([]byte, int, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	if disableRedirects {
		client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}
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

	if insecure {
		log.Println("-insecure flag, disable TLS verification")
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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
