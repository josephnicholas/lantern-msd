package core

import (
	"fmt"
	"github.com/vbauerster/mpb"
	"github.com/vbauerster/mpb/decor"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
)

type Downloader struct {
	File           FileDetails
	ChunkSize      int64
	NumberOfChunks int64
}

func (d *Downloader) Execute() {
	fileDetails := d.File.GetFileDetails()
	numberOfChunks := d.calculateConcurrentWorker()

	wg := sync.WaitGroup{}
	chunks := make([]chan []byte, numberOfChunks)
	pb := mpb.New(mpb.WithWaitGroup(&wg))
	bars := make([]*mpb.Bar, numberOfChunks)

	for i := int64(0); i < numberOfChunks; i++ {
		chunks[i] = make(chan []byte)

		startByte, endByte, currentChunkSize := d.calculateChunk(i, numberOfChunks, fileDetails)
		if currentChunkSize <= 0 {
			continue
		}

		// extract this to a function
		counter := int(i % fileDetails.Urls.GetSize())
		fileUrl := fileDetails.Urls.GetUrls()[counter]
		parsedURL, err := url.Parse(fileUrl)
		if err != nil {
			log.Fatal(err)
		}

		barSize := currentChunkSize
		bars[i] = pb.AddBar(barSize,
			mpb.PrependDecorators(
				decor.CountersKibiByte("% .2f / % .2f"),
				decor.Name(fmt.Sprintf(" [%s]", parsedURL.Host)),
			),
			mpb.AppendDecorators(
				decor.EwmaETA(decor.ET_STYLE_MMSS, float64(d.ChunkSize)/2048),
				decor.Name(" ] "),
				decor.AverageSpeed(decor.UnitKiB, "% .2f"),
			),
		)

		wg.Add(1)
		go d.downloadWorker(fileUrl, startByte, endByte, &wg, bars[i], &chunks[i])
	}

	go func() {
		wg.Wait()
		for i := range chunks {
			close(chunks[i])
		}
	}()

	const flag = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	WriteToFile(flag, fileDetails.Name, chunks)
}

func (d *Downloader) calculateChunk(i int64, numberOfChunks int64, fileDetails *FileDetails) (int64, int64, int64) {
	startByte := i * d.ChunkSize
	endByte := startByte + d.ChunkSize - 1
	if i+1 == numberOfChunks || endByte > fileDetails.Size-1 {
		endByte = fileDetails.Size - 1
	}

	currentChunkSize := endByte - startByte + 1
	return startByte, endByte, currentChunkSize
}

func (d *Downloader) downloadWorker(fileUrl string, startByte int64, endByte int64, wg *sync.WaitGroup, bar *mpb.Bar, chunk *chan []byte) {
	defer wg.Done()
	req, err := http.NewRequest("GET", fileUrl, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", startByte, endByte))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to do response: %v", err)
	}
	defer resp.Body.Close()

	reader := bar.ProxyReader(resp.Body)

	data, err := io.ReadAll(reader)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	*chunk <- data
}

func (d *Downloader) calculateConcurrentWorker() int64 {
	if d.NumberOfChunks != 0 {
		return d.NumberOfChunks
	}

	return d.File.Size / d.ChunkSize
}
