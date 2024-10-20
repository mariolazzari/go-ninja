package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.99,
	}

	// full map
	fmt.Println(menu)

	// return value by key
	fmt.Println(menu["pie"])

	// loop
	for key, value := range menu {
		fmt.Printf("key %v, value %v\n", key, value)
	}

	// ints as key type
	phones := map[int]string{
		123456: "Mario",
		123457: "Mariarosa",
		123458: "Maria",
	}
	fmt.Println(phones)
	fmt.Println(phones[123456])

	// update item
	phones[123456] = "Mario Lazzari"
	fmt.Println(phones)

}
