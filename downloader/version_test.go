package downloader

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateShaAndCompare(t *testing.T) {
	t.Run("when file exists", func(t *testing.T) {
		defer os.Remove(serverVersionFile)
		file, err := os.OpenFile(serverVersionFile, os.O_RDWR|os.O_CREATE, 0644)
		file.Write([]byte("3.2.1"))
		assert.NoError(t, err)

		downloader := New("some-fancy.web")

		assert.Equal(t, []int{3, 2, 1}, downloader.version.numeric)
		assert.Equal(t, "3.2.1", downloader.version.raw)
	})

	t.Run("when file does not exist", func(t *testing.T) {
		defer os.Remove(serverVersionFile)
		downloader := New("some-fancy.web")

		assert.Equal(t, []int{0, 0, 0}, downloader.version.numeric)
		assert.Equal(t, "0.0.0", downloader.version.raw)

		_, err := os.Open(serverVersionFile)
		assert.NoError(t, err)
	})
}
