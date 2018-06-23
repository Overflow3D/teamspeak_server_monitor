package downloader

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type version struct {
	raw     string
	numeric []int
}

const serverVersionFile = ".server_version"

// IsNewVersionAvailable check if recentVersion is the same version as server version
func (d *Downloader) serverVersion() {
	rawFile, err := ioutil.ReadFile(serverVersionFile) // just pass the file name
	if err != nil {
		makeFileAssignVersion(d)
		d.serverVersion()
		return
	}
	serverVersion := string(rawFile)

	d.version.raw = serverVersion

	versionString := strings.Split(serverVersion, ".")
	numberCount := len(versionString)
	numericVersion := make([]int, 0, numberCount)

	for i := 0; i < numberCount; i++ {
		number, _ := strconv.Atoi(versionString[i])
		numericVersion = append(numericVersion, number)
	}

	d.version.numeric = numericVersion
}

func accessFile() (*os.File, error) {
	file, err := os.OpenFile(serverVersionFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func getServerVersion(file *os.File) string {
	bytes := make([]byte, 10) // since we will store only char like 3.2.2, size of 10 should be enough
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func makeFileAssignVersion(d *Downloader) {
	file, err := os.OpenFile(serverVersionFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// add reading version from current teamspeak3 server dir
	file.Write([]byte("0.0.0"))
	d.version.numeric = make([]int, 0, 4)
	d.version.raw = "0.0.0"
}
