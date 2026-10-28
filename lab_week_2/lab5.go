package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	var text string
	fmt.Print("Enter a text: ")
	fmt.Scanln(&text)

	data := []byte(text)

	fmt.Print("\nBinary: ")
	for _, b := range data {
		fmt.Printf("\n%08 ", b)
	}

	fmt.Printf("\nHex		:", hex.EncodeToString(data))

	fmt.Println("\nBase64		:", base64.StdEncoding.EncodeToString(data))
}
