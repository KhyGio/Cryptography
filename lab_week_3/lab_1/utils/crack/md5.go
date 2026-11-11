package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
)

func CrackMD5(targetHash string, wordlistPath string) (string, error) {
	file, err := os.Open(wordlistPath)
	if err != nil {
		return "", fmt.Errorf("could not open wordlist file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	fmt.Printf("[*] Searching for hash: %s\n", targetHash)

	for scanner.Scan() {
		word := scanner.Text()

		hasher := md5.New()
		hasher.Write([]byte(word))
		currentHash := hex.EncodeToString(hasher.Sum(nil))

		if currentHash == targetHash {
			return word, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading wordlist: %w", err)
	}

	return "", fmt.Errorf("password not found in the provided wordlist")
}
