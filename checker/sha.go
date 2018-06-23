package checker

import (
	"crypto/sha256"
	"fmt"
)

// CalculateShaAndCompare transform file into sha256 and compare it with current website sha256
func CalculateShaAndCompare(downloadedFile []byte, websiteSha string) bool {
	return calculateSha(downloadedFile) == websiteSha
}

func calculateSha(file []byte) string {
	return fmt.Sprintf("%x", sha256.Sum256(file))
}
