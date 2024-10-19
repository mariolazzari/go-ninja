package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Println("x =", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}

	names := []string{"Mario", "Mariarosa", "Roberto", "Maria"}
	for i := 0; i < len(names); i++ {
		fmt.Println("name", i, names[i])
	}

	// range
	for index, value := range names {
		fmt.Printf("Index %v = %q\n", index, value)
	}

	// no index
	for _, value := range names {
		fmt.Printf("Value %q\n", value)
	}

	// no value
	for index := range names {
		fmt.Printf("Index %v\n", index)
	}
}
