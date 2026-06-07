package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func main() {

	// Create a new array containing 3 hobbies
	hobbies := [3]string{"coding", "football", "swimming"}
	fmt.Println("My hobbies:", hobbies)

	// Access the first hobby and print it
	fmt.Println("First hobby:", hobbies[0])

	// Access the second and third hobbies and print them
	fmt.Println("Rest hobbies:", hobbies[1:3])

	// Create a new slice containing the first two hobbies
	mainHobbies := hobbies[:2]
	fmt.Println("\nMain hobbies:", mainHobbies)

	// Reslice the slice to include all hobbies
	fmt.Println("Size of the main hobbies:", len(mainHobbies), "and capacity:", cap(mainHobbies))

	allHobbies := mainHobbies[:3]
	fmt.Println("All hobbies:", allHobbies)

	// Create a new dynamic array of goals
	myGoals := []string{"learn Go", "travel the world", "start a business"}
	fmt.Println("\nMy goals:", myGoals)

	// Append a new goal to the dynamic array
	myGoals = append(myGoals, "write a book")
	fmt.Println("Updated goals:", myGoals)

	// Remove the second goal from the dynamic array
	myGoals = append(myGoals[:1], myGoals[2:]...) // why ...? Because it unpacks the slice
	fmt.Println("Goals after removal:", myGoals)

	// Create a dynamic array of structs
	products := []Product{
		{Name: "Laptop", Price: 999.99},
		{Name: "Smartphone", Price: 499.99},
	}
	fmt.Println("\nProducts:", products)

	// Append a new product to the dynamic array of structs
	products = append(products, Product{Name: "Tablet", Price: 299.99})
	fmt.Println("Updated products:", products)
}
