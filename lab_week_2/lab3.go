package main

import "fmt"

func myAND(a, b int) {
	fmt.Printf("AND    : %2d & %2d = %2d   (%08b & %08b = %08b)\n", a, b, a&b, a, b, a&b)
}

func myOR(a, b int) {
	fmt.Printf("OR     : %2d | %2d = %2d   (%08b | %08b = %08b)\n", a, b, a|b, a, b, a|b)
}

func myXOR(a, b int) {
	fmt.Printf("XOR    : %2d ^ %2d = %2d   (%08b ^ %08b = %08b)\n", a, b, a^b, a, b, a^b)
}

func myNOT(a, b int) {
	fmt.Printf("NOT    : %2d = %3d       (^%08b = ^%08b)\n", a, ^a, a, ^a)
	fmt.Printf("NOT    : %2d = %3d       (^%08b = ^%08b)\n", b, ^b, b, ^b)
}

func main() {
	var a int
	var b int

	fmt.Printf("Enter a: ")
	fmt.Scan(&a)

	fmt.Printf("Enter b: ")
	fmt.Scan(&b)

	myXOR(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
}

