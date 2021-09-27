package main

import "fmt"


func main() {
	var method string
	fmt.Println("What method do you want this calculator in?\n\n(-, +, *, /)")
	fmt.Scanln(&method)
	if (method) == "+" {
		var num1 int
		var num2 int
		fmt.Println("num 1:")
		fmt.Scanln(&num1)
		fmt.Println("num 2:")
		fmt.Scanln(&num2)
		fmt.Println(num1 + num2)
	}

	if (method) == "-" {
		var num1 int
		var num2 int
		fmt.Println("num 1:")
		fmt.Scanln(&num1)
		fmt.Println("num 2:")
		fmt.Scanln(&num2)
		fmt.Println(num1 - num2)
	}

	if (method) == "*" {
		var num1 int
		var num2 int
		fmt.Println("num 1:")
		fmt.Scanln(&num1)
		fmt.Println("num 2:")
		fmt.Scanln(&num2)
		fmt.Println(num1 * num2)
	}

	if (method) == "/" {
		var num1 int
		var num2 int
		fmt.Println("num 1:")
		fmt.Scanln(&num1)
		fmt.Println("num 2:")
		fmt.Scanln(&num2)
		fmt.Println(num1 / num2)
}
}
