package hashs

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func CheckIntegrity(filename string, expectedHashes map[string]string) (bool, error) {
	originalDir, _ := os.Getwd()
	f, err := os.Open(originalDir + "/banners/" + filename + ".txt")
	if err != nil {
		fmt.Print(err)
		return false, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	actualHash := fmt.Sprintf("%x", h.Sum(nil))
	expectedHash, ok := expectedHashes[filename+".txt"]
	if !ok {
		return false, fmt.Errorf("unknown banner")
	}

	return actualHash == expectedHash, nil
}
