package collector

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Collector struct {
	client *http.Client
}

func NewCollector(duration time.Duration) *Collector {
	return &Collector{
		client: &http.Client{
			Timeout: duration,
		},
	}
}

func (c *Collector) FetchData(urls chan string, result chan string) {
	var wg sync.WaitGroup
	for v := range urls {
		wg.Add(1)
		go func(v string) {
			defer wg.Done()
			start := time.Now()

			resp, err := c.fetch(v)
			if err != nil {
				result <- fmt.Sprintf("Url: %s Error: %s", v, err.Error())
				return
			}
			result <- fmt.Sprintf("Url: %s Time: %s ContentLength: %d", v, time.Since(start), resp.ContentLength)
		}(v)
	}
	wg.Wait()
}

func (c *Collector) fetch(u string) (*http.Response, error) {
	_, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Get(u)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
