package core

import (
	"errors"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

type FileDetails struct {
	Urls UrlCollector
	Name string
	Size int64
	Etag string
}

func NewFileDetails(urls UrlCollector) (*FileDetails, error) {
	// This assumes all other urls points to the same file
	urls.FilterValidUrl()
	fileUrl, err := urls.GetFirstUrl()
	if err != nil {
		return nil, errors.New("no valid url found")
	}

	resp, err := http.Head(fileUrl)
	if err != nil {
		return nil, errors.New("failed to get file details")
	}
	defer resp.Body.Close()

	parsedURL, err := url.Parse(fileUrl)
	if err != nil {
		return nil, errors.New("failed to parse url")
	}

	fileName := path.Base(parsedURL.Path)
	fileSize, err := strconv.ParseInt(resp.Header.Get("Content-Length"), 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse content-length")
	}

	f := &FileDetails{
		Urls: urls,
		Name: fileName,
		Size: fileSize,
		Etag: resp.Header.Get("Etag"),
	}

	return f, nil
}

func (details *FileDetails) GetFileDetails() *FileDetails {
	return details
}
