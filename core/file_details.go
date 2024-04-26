package core

import (
	"log"
	"net/http"
	"net/url"
	"path"
)

type FileDetails struct {
	Urls UrlCollector
	Name string
	Size int64
	Etag string
}

func NewFileDetails(urls UrlCollector) *FileDetails {
	// This assumes all other urls points to the same file
	fileUrl := urls.GetFirstUrl()
	resp, err := http.Head(fileUrl)
	if err != nil {
		log.Fatalf("Failed to get file details: %v", err)
	}
	defer resp.Body.Close()

	parsedURL, err := url.Parse(fileUrl)
	if err != nil {
		log.Fatal(err)
	}

	fileName := path.Base(parsedURL.Path)
	f := &FileDetails{
		Urls: urls,
		Name: fileName,
		Size: resp.ContentLength,
		Etag: resp.Header.Get("Etag"),
	}

	return f
}

func (details *FileDetails) GetFileDetails() *FileDetails {
	return details
}

// maybe use a tuple to return
