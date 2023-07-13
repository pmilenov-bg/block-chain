package services

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestHashFromString(t *testing.T) {
	message := "Hello, World!"
	expectedHash := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"

	hash := HashFromString(message)
	hashString := hex.EncodeToString(hash)

	if hashString != expectedHash {
		t.Errorf("Expected hash '%s', but got '%s'", expectedHash, hashString)
	}
}

func TestHashBlock(t *testing.T) {
	data := "Block Data"
	prevHash := []byte("Previous Hash")
	expectedHash := "9263fd4e2b0c98187076884673346392313a8f7810838dd21a62ff096fcbdd11"

	hash := HashBlock(data, prevHash)
	hashString := hex.EncodeToString(hash)

	if hashString != expectedHash {
		t.Errorf("Expected hash '%s', but got '%s'", expectedHash, hashString)
	}
}

func TestHashBlockWithEmptyPrevHash(t *testing.T) {
	data := "Block Data"
	prevHash := []byte{}

	hash := HashBlock(data, prevHash)

	if len(hash) != sha256.Size {
		t.Errorf("Expected hash size of %d, but got %d", sha256.Size, len(hash))
	}
}
