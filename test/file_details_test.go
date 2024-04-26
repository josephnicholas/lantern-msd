package test

import (
	"github.com/nbio/st"
	"lantern-msd/core"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConstructNewDetails(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "104857600")
		w.Header().Set("Etag", "123456789")
	}))

	urls := core.NewUrlCollector(server.URL)
	details, _ := core.NewFileDetails(*urls)

	st.Expect(t, details.Urls.GetSize(), int64(1))
	st.Expect(t, details.Size, int64(104857600))
	st.Expect(t, details.Etag, "123456789")
}

func TestConstructDetailsInvalidURL(t *testing.T) {
	urls := core.NewUrlCollector("hello_world")
	_, err := core.NewFileDetails(*urls)
	st.Expect(t, err.Error(), "no valid url found")
}

func TestConstructDetailsInvalidLength(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "invalid")
	}))

	urls := core.NewUrlCollector(server.URL)
	_, err := core.NewFileDetails(*urls)
	st.Expect(t, err.Error(), "no valid url found")
}

func TestConstructDetailsMethodNotAllowed(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}))

	urls := core.NewUrlCollector(server.URL)
	_, err := core.NewFileDetails(*urls)
	st.Expect(t, err.Error(), "no valid url found")
}
