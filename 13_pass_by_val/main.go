package main

import "fmt"

func updateName(n string) {
	n = "Mario"
}

func updateMenu(m map[string]float64) {
	m["coffee"] = 1.99
}

func main() {
	// group A type
	n := "mario"
	updateName(n)
	fmt.Println(n)

	// group b type
	menu := map[string]float64{
		"pie":       3.99,
		"ice cream": 2.99,
	}
	updateMenu(menu)
	fmt.Println(menu)
}
