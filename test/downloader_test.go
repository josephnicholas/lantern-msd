package test

import (
	"github.com/nbio/st"
	"lantern-msd/core"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestDownloadWorker(t *testing.T) {
	expected := []byte{1, 2, 3, 4, 5}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
	}))

	start := int64(0)
	end := int64(4)

	d := core.Downloader{
		ChunkSize:      1,
		NumberOfChunks: 1,
	}

	chunk := make(chan []byte)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go d.DownloadWorker(server.URL, start, end, &wg, nil, &chunk)

	data := <-chunk
	st.Expect(t, data, expected)
}

func TestDownloadChunkTwo(t *testing.T) {
	expected := []byte{1, 2, 3, 4, 5}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
	}))

	start := int64(0)
	end := int64(4)

	d := core.Downloader{
		ChunkSize:      3,
		NumberOfChunks: 5,
	}

	chunk := make(chan []byte)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go d.DownloadWorker(server.URL, start, end, &wg, nil, &chunk)

	data := <-chunk
	st.Expect(t, data, expected)
}

func TestDownloadWorkerFailed(t *testing.T) {
	servers := []*httptest.Server{
		httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})),
		httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotFound)
		})),
		httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
		})),
		httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusForbidden)
		})),
	}

	d := core.Downloader{
		ChunkSize:      1,
		NumberOfChunks: 1,
	}

	chunk := make(chan []byte)
	wg := sync.WaitGroup{}

	for _, server := range servers {
		wg.Add(1)
		go d.DownloadWorker(server.URL, 0, 4, &wg, nil, &chunk)

		data := <-chunk
		st.Expect(t, len(data), 0)
	}
}

func TestDownloadWorkerReceivedPartialContent(t *testing.T) {
	expected := []byte{1, 2, 3, 4, 5}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(expected)
		w.WriteHeader(http.StatusPartialContent)
	}))

	start := int64(0)
	end := int64(4)

	d := core.Downloader{
		ChunkSize:      1,
		NumberOfChunks: 1,
	}

	chunk := make(chan []byte)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go d.DownloadWorker(server.URL, start, end, &wg, nil, &chunk)

	data := <-chunk
	st.Expect(t, data, expected)
}
