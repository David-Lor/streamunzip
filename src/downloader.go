package main

import (
	"fmt"
	"net/http"
)

type Downloader struct {
	DownloadURL string
	OutputPath  string

	response *http.Response
}

func (d *Downloader) Download() error {
	defer d.close()
	if err := d.httpGet(); err != nil {
		return err
	}

	if err := d.readZipInteractively(); err != nil {
		return err
	}

	return nil
}

func (d *Downloader) close() {
	if d.response != nil {
		d.response.Body.Close()
	}
}

func (d *Downloader) httpGet() error {
	var err error
	url := d.DownloadURL
	d.response, err = http.Get(url)
	if err != nil {
		return fmt.Errorf("failed GET %s: %v", url, err)
	}

	statusCode := d.response.StatusCode
	if statusCode < 200 || statusCode >= 300 {
		return fmt.Errorf("GET %s returned statuscode %d", url, statusCode)
	}

	contentType := d.response.Header.Get("Content-Type")
	if contentType == "" {
		return fmt.Errorf("GET %s returned no Content-Type", url)
	}
	if contentType != ZipMimeType {
		return fmt.Errorf("GET %s returned invalid Content-Type=%s", url, contentType)
	}
	return nil
}

func (d *Downloader) readZipInteractively() error {
	extractor := ZipExtractor{
		ZipStream:  d.response.Body,
		OutputPath: d.OutputPath,
	}
	return extractor.ExtractInteractively()
}
