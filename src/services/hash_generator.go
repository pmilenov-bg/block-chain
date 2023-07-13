package services

import (
	"bytes"
	"crypto/sha256"
	// "encoding/hex"
)

func HashFromString(message string) []byte {
	h := sha256.New()
	h.Write([]byte(message))
	// sha1_hash := hex.EncodeToString(h.Sum(nil))

	return h.Sum(nil)
}

func HashBlock(data string, prevHash []byte) []byte {
	byteSlice := []byte(data)
	headers := bytes.Join([][]byte{prevHash, byteSlice}, []byte{})
	hash := sha256.Sum256(headers)

	return hash[:]
}
