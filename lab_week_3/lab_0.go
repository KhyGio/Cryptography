package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func hashString(algorithm string, text string) string {
	var hashBytes []byte

	switch algorithm {
	case "MD5":
		h := md5.Sum([]byte(text))
		hashBytes = h[:]

	case "SHA1":
		h := sha1.Sum([]byte(text))
		hashBytes = h[:]

	case "SHA256":
		h := sha256.Sum256([]byte(text))
		hashBytes = h[:]

	case "SHA512":
		h := sha512.Sum512([]byte(text))
		hashBytes = h[:]

	case "SHA3":
		h := sha256.Sum256([]byte(text))
		hashBytes = h[:]

	}
	return hex.EncodeToString(hashBytes)
}

func proofMe(input1, input2 string) {
	fmt.Println("==========Name + Hashing Program==========")
	fmt.Printf("input value 1: %s\n", input1)
	fmt.Printf("input value 2: %s\n", input2)

	algorithms := []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3"}

	for _, algo := range algorithms {
		outputA := hashString(algo, input1)
		outputB := hashString(algo, input2)

		fmt.Printf("Hash (%s):\n", algo)
		fmt.Printf("Output A = %s\n", outputA)
		fmt.Printf("Output B = %s\n", outputB)

		if outputA == outputB {
			fmt.Println(" => Match!")
		} else {
			fmt.Println(" => No Match!")
		}
		fmt.Println()
	}
}

func main() {
	var input1 string
	var input2 string

	fmt.Printf("Enter Text: ")
	fmt.Scan(&input1)

	fmt.Printf("Enter Text: ")
	fmt.Scan(&input2)

	proofMe(input1, input2)
}
