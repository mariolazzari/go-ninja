package main

import "fmt"

func main() {
	const maxAge uint8 = 50
	const age uint8 = 45

	fmt.Println(age <= maxAge)
	fmt.Println(age >= maxAge)
	fmt.Println(age == maxAge)
	fmt.Println(age != maxAge)

	// if
	if age < 30 {
		fmt.Println("Age less than 30")
	} else if age < 40 {
		fmt.Println("Age less than 40")
	} else {
		fmt.Println("Age more than 45")
	}

	// if nested in loop
	names := []string{"Mario", "Mariarosa", "Roberto", "Maria"}
	for index, value := range names {
		if index == 1 {
			fmt.Printf("Continue at index %v: %q\n", index, value)
			continue
		}

		if index > 2 {
			fmt.Printf("Breaking at index %v: %q\n", index, value)
			break
		}

		fmt.Printf("Index %v: %q\n", index, value)
	}
}
