package downloader

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	node, err := ioutil.ReadFile("../helpers/teamspeak_node.html")
	assert.NoError(t, err)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(node)))
	assert.NoError(t, err)

	mainNode := doc.Find(selectors["root"]).First()

	updateInfo, err := parseNodesInformation(mainNode)
	assert.NoError(t, err)

	assert.Equal(t, "http://dl.4players.de/ts/releases/3.2.0/teamspeak3-server_linux_amd64-3.2.0.tar.bz2", updateInfo["url"])
	assert.Equal(t, "3.2.0", updateInfo["version"])
	assert.Equal(t, "f1e267334e8863342e8eb90ae22203b761b54d9d4400a25ed1fd34fce2187f57", updateInfo["sha"])
}
