package hashGenerator

import (
    "crypto/sha256"
    "encoding/hex"
)

func HashString1(message string)  string{

    h := sha256.New()
    h.Write([]byte(message))
    sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
    // fmt.Println(message, sha1_hash)

}

func HashString2(message string, prevHash string) string {

	h := sha256.New()
    h.Write([]byte(message + prevHash))
    sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}