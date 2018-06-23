package downloader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

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
	err = d.getNewVersion(updateInfo)
	if err != nil {
		fmt.Println(err)
	}
	return d.quit, nil
}

func (d *Downloader) getNewVersion(updateInfo map[string]string) error {
	downloadedFileName := fmt.Sprintf("new_%s.tar.bz2", updateInfo["version"])

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

	log.Printf("Downloaded %d byte file.\n", numBytesWritten)
	return nil

}

func createNewFile(name string) (*os.File, error) {
	newFile, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return newFile, nil
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
