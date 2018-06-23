package downloader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Downloader ...
type Downloader struct {
	url     string
	version *version
	quit    chan struct{}
}

const (
	emptyFileSize     = 0
	newFileNameFormat = "new_%s.tar.bz2"
)

// New creates download struct
func New(url string) *Downloader {
	downloader := &Downloader{
		url:     url,
		version: &version{},
	}
	downloader.serverVersion()
	return downloader
}

// StartUpdater starts ongoing check for new version
func (d *Downloader) StartUpdater() (chan struct{}, error) {
	go d.updateCheckScheduler()
	return d.quit, nil
}

func (d *Downloader) updateCheckScheduler() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			// add better logs
			updateInfo, err := d.gatherInformation()
			log.Println(err)
			err = d.fetchRecentServerFiles(updateInfo)
			log.Println(err)
		}
	}
}

func (d *Downloader) fetchRecentServerFiles(updateInfo map[string]string) error {
	version := updateInfo["version"]

	downloadedFileName := fmt.Sprintf(newFileNameFormat, version)
	if d.version.raw == version {
		return nil
	}

	newFile, err := createNewFile(downloadedFileName)
	if err != nil {
		return err
	}
	defer newFile.Close()

	downloadedFile, err := downloadFile(updateInfo["url"], updateInfo["sha"])
	if err != nil {
		return err
	}
	defer downloadedFile.Close()

	numBytesWritten, err := io.Copy(newFile, downloadedFile)
	if err != nil {
		return err
	}

	defer func() {
		// create log info here.
		if unsuccessfulDownload(numBytesWritten) {
			os.Remove(downloadedFileName)
		}
	}()

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	return nil
}

func downloadFile(url string, sha string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if correct := isShaCorrect(body, sha); !correct {
		return nil, fmt.Errorf("incorrect sha256, file might be corrupted")
	}

	// since we drained body with sha check we need to regenerate it back
	return ioutil.NopCloser(bytes.NewBuffer(body)), err
}

func unsuccessfulDownload(fileSize int64) bool {
	return fileSize == emptyFileSize
}
