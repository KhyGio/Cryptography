package crack

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func CrackSHA512(targetHash string, wordlist []string) (string, int) {
	for i, word := range wordlist {

		hasher := sha512.New()
		hasher.Write([]byte(word))
		currentHash := hex.EncodeToString(hasher.Sum(nil))

		fmt.Printf("Attempt %d: Testing word '%s' -> Hash: %s\n", i+1, word, currentHash)

		if currentHash == targetHash {
			return word, i + 1
		}
	}
	return "", -1
}
