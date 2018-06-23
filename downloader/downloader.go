package downloader

import "fmt"

// Downloader ...
type Downloader struct {
	url            string
	currentVersion string
	quit           chan struct{}
}

// New creates download struct
func New(url string) *Downloader {
	return &Downloader{url: url}
}

// StartUpdater starts ongoing check for new version
func (d *Downloader) StartUpdater() (chan struct{}, error) {
	updateInfo, err := d.gatherInformation()
	if err != nil {
		return nil, err
	}
	fmt.Println(updateInfo)
	return d.quit, nil
}
