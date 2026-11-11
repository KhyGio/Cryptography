package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lab_3/utils/crack"
)

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	file, err := os.Open("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordlist = append(wordlist, strings.TrimSpace(scanner.Text()))
	}

	fmt.Println("Starting SHA-512 password cracking process...")

	password, attempts := crack.CrackSHA512(targetHash, wordlist)

	if password != "" {
		fmt.Printf("\n[SUCCESS] Password found: %s\n", password)
		fmt.Printf("Total attempts: %d\n", attempts)
	} else {
		fmt.Println("\n[FAILED] Password not found in the wordlist.")
	}
}
