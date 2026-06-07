package main

import "fmt"

func recursion() {
	fmt.Print("Enter a number to find it's factorial: ")
	var num int
	fmt.Scanln(&num)
	fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return factorial(num-1) * num
}
