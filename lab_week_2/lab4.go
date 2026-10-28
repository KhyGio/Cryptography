package main

import "fmt"

func operator(a, b, o int) {
	switch o {
	case 1:
		fmt.Println("Result: ", a+b)
	case 2:
		fmt.Println("Result: ", a-b)
	case 3:
		fmt.Println("Result: ", a*b)
	case 4:
		if b != 0 {
			fmt.Println("Result: ", a/b)
		} else {
			fmt.Println("[!] Cannot Divide by 0")
		}
	case 5:
		if b != 0 {
			fmt.Println("Result: ", a%b)
		} else {
			fmt.Println("[!] Cannot Divide by 0")
		}
	default:
		fmt.Println("[!] Invalid Operator.")
	}
}

func main() {
	for {
		var o int
		var a int
		var b int

		fmt.Println(`
===== Mini Calculator =====
1) Add  2) Sub  3) Mul
4) Div  5) Mod  6) Exit`)

		fmt.Printf("Enter a: ")
		fmt.Scan(&a)

		fmt.Printf("Enter b: ")
		fmt.Scan(&b)

		fmt.Print("Operator: ")
		fmt.Scan(&o)
		if o == 6 {
			fmt.Println("Exit !")
			break
		}

		operator(a, b, o)
	}
}
