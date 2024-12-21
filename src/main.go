package main

import (
	"log"
)

func amain() error {
	InitStdinReader()

	downloadURL := ReadInput("Enter ZIP URL to download")
	downloadPath := ReadInput("Enter download path")

	downloader := Downloader{
		DownloadURL: downloadURL,
		OutputPath:  downloadPath,
	}
	return downloader.Download()
}

func main() {
	if err := amain(); err != nil {
		log.Fatal(err)
	}
}
