package checker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateShaAndCompare(t *testing.T) {
	exampleFile := []byte("123")

	t.Run("when matches", func(t *testing.T) {
		isMatched := CalculateShaAndCompare(exampleFile, "a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3")
		assert.True(t, isMatched)
	})

	t.Run("when not match", func(t *testing.T) {
		isMatched := CalculateShaAndCompare(exampleFile, "imma not sha")
		assert.False(t, isMatched)
	})
}
