package crack

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

func CrackSHA1(targetHash string, wordlist []string) (string, int) {
	for i, word := range wordlist {
		hasher := sha1.New()
		hasher.Write([]byte(word))
		hash := hex.EncodeToString(hasher.Sum(nil))

		fmt.Printf("Trying: %s -> %s\n", word, hash)

		if hash == targetHash {
			return word, i + 1
		}
	}
	return "", -1
}
