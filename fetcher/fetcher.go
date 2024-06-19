package fetcher

import (
	"io"
	"net/http"
	"time"
)

func Fetch() ([]byte, error) {
	dummyUrl, timeout1 := "https://go.dev/doc/", 10*time.Second
	client := http.Client{
		Timeout: timeout1,
	}
	resp, err := client.Get(dummyUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
