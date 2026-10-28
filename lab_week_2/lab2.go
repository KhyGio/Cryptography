package main

import "fmt"

func main() {
	var x int
	var y int

	fmt.Print("Enter x: ")
	fmt.Scan(&x)

	fmt.Print("Enter y: ")
	fmt.Scan(&y)

	fmt.Println("x > y")
	if x > y {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	fmt.Println("x & y: Positive")
	if x > 0 && y > 0 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	fmt.Println("x != y")
	if x != y {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
