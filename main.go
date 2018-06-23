package main

import "github.com/Overflow3D/teamspot_monitor/checker"

const (
	licenseAccept = "TS3SERVER_LICENSE=accept"
)

func main() {
	// dl := downloader.New("https://www.teamspeak.com/en/downloads#server")
	// dl.StartUpdater()
	checker.IsNewVersionAvailable("")
}
