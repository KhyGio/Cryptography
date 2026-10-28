package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var x int
	var y int

	fmt.Print("Enter x: ")
	fmt.Scan(&x)

	fmt.Print("Enter y: ")
	fmt.Scan(&y)

	r := 10
	r += x
	fmt.Printf("Result: %d\n", r)

	r -= y
	fmt.Printf("Result: %d\n", r)

	r *= x
	fmt.Printf("Result: %d\n", r)

	r /= y
	fmt.Printf("Result: %d\n", r)

	r %= x
	fmt.Printf("Result: %d\n", r)
}