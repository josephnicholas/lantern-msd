package core

import (
	"fmt"
	"log"
	"os"
)

func WriteToFile(flag int, fileName string, chunks []chan []byte) {
	f, err := os.OpenFile(fmt.Sprintf("%s", fileName), flag, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer f.Close()

	for i := range chunks {
		chunk := <-chunks[i]
		f.Write(chunk)
	}
}
