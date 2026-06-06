package main

import "fmt"

func main() {
	value := 37

	var valuePointer *int = &value

	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Value Pointer: %p\n", valuePointer)
	fmt.Printf("Dereferenced Value Pointer: %d\n", *valuePointer)

	fmt.Printf("Value after change: %d\n", changeValue(valuePointer))
	fmt.Printf("Value after change (original): %d\n", value)
}

func changeValue(val *int) int {
	*val = *val - 10
	return *val
}
