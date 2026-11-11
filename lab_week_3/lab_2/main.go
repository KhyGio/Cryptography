package main

import (
	"bufio"
	"fmt"
	"lab_2/utils/crack"
	"os"
	"strings"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

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

	fmt.Println("Starting crack...")
	password, attempts := crack.CrackSHA1(targetHash, wordlist)

	if password != "" {
		fmt.Printf("\n[SUCCESS] Password found: %s\n", password)
		fmt.Printf("Attempts: %d\n", attempts)
	} else {
		fmt.Println("\n[FAILED] Password not found in wordlist.")
	}
}
