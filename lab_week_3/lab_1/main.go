package main

import (
	"fmt"
	"lab_1/utils/crack"
	"os"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	wordlistPath := "nord_vpn.txt"

	fmt.Println("========== Password Cracking Program ==========")

	password, err := crack.CrackMD5(targetHash, wordlistPath)
	if err != nil {
		fmt.Printf("[-] Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("[+] SUCCESS! Password found: %s\n", password)
	fmt.Println("===============================================")
}
