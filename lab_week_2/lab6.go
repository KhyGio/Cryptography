package main

import (
	"fmt"
)

func xorEncrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}

func main() {
	// 1. Input
	var plaintext string
	var keyChar string

	fmt.Print("Enter plaintext : ")
	fmt.Scanln(&plaintext)
	fmt.Print("Enter key (1 char): ")
	fmt.Scanln(&keyChar)

	// 2. Make sure the key is exactly one byte
	if len(keyChar) != 1 {
		fmt.Println("Key must be exactly one character.")
		return
	}
	key := keyChar[0]

	// 3. Encrypt
	ciphertext := xorEncrypt(plaintext, key)

	// 4. Decrypt (same function, same key!)
	recovered := xorEncrypt(ciphertext, key)

	// 5. Show results
	fmt.Println("\n=== Results ===")
	fmt.Printf("Plaintext  : %q\n", plaintext)
	fmt.Printf("Key        : %q (byte value %d)\n", string(key), key)
	fmt.Printf("Ciphertext : %q\n", ciphertext)
	fmt.Printf("Ciphertext (hex): %x\n", ciphertext)
	fmt.Printf("Decrypted  : %q\n", recovered)

	if recovered == plaintext {
		fmt.Println("\n[*] Decryption successful — matches original!")
	} else {
		fmt.Println("\n[!] Something went wrong.")
	}
}
