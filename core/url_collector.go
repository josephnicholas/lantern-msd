package core

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

func (c *UrlCollector) GetFirstUrl() string {
	return c.urls[0]
}

func (c *UrlCollector) GetUrls() []string {
	return c.urls
}

func (c *UrlCollector) GetSize() int64 {
	return int64(len(c.urls))
}
