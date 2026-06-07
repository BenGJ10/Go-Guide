package main

import "fmt"

func closures() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Initial numbers: ", numbers)

	double := createTransformer(2)
	triple := createTransformer(3)

	doubledNumbers := transformNums(numbers, double)
	fmt.Println("Numbers after doubling: ", doubledNumbers)

	tripledNumbers := transformNums(numbers, triple)
	fmt.Println("Numbers after tripling: ", tripledNumbers)
}

// A closure is a function that captures variables from its surrounding scope.
// In Go, closures are created when you define a function inside another function and the inner function references variables from the outer function.

// Here we are creating a closure that multiplies a number by a given multiplier
// The createTransformer function returns a closure that captures the multiplier variable
// The returned function can be used to transform numbers by multiplying them with the captured multiplier
func createTransformer(multiplier int) func(int) int {
	return func(num int) int {
		return num * multiplier
	}
}

func transformNums(nums []int, transform func(int) int) []int {
	result := []int{}
	for _, num := range nums {
		result = append(result, transform(num))
	}
	return result
}
