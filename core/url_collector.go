package core

import (
	"errors"
	"net/http"
)

type UrlCollector struct {
	urls []string
}

func NewUrlCollector(urls ...string) *UrlCollector {
	urlContainer := make([]string, 0)
	for _, url := range urls {
		urlContainer = append(urlContainer, url)
	}

	return &UrlCollector{urls: urlContainer}
}

func (c *UrlCollector) GetFirstUrl() (string, error) {
	if len(c.urls) == 0 {
		return "", errors.New("url list is empty")
	}
	return c.urls[0], nil
}

func (c *UrlCollector) GetUrls() []string {
	return c.urls
}

func (c *UrlCollector) GetSize() int64 {
	return int64(len(c.urls))
}

func (c *UrlCollector) FilterValidUrl() {
	var validUrls []string
	for _, urlToCheck := range c.GetUrls() {
		resp, err := http.Head(urlToCheck)
		if err != nil || resp.StatusCode != http.StatusOK {
			continue
		}
		validUrls = append(validUrls, urlToCheck)
	}

	c.urls = validUrls
}
