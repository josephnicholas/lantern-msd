package main

import "lantern-msd/core"

func main() {
	urls := core.NewUrlCollector(
		"https://mirror.sitsa.com.ar/ubuntu-releases/noble/ubuntu-24.04-desktop-amd64.iso",
		"https://mirror.xenyth.net/ubuntu-releases/noble/ubuntu-24.04-desktop-amd64.iso")

	f := core.NewFileDetails(*urls)

	d := core.Downloader{
		File:           *f,
		NumberOfChunks: 10,
		ChunkSize:      1024 * 1024 * 200}

	d.Execute()
}
