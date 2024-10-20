package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// add new bill
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"pie":    3.99,
			"coffee": 2.88,
		},
		tip: 0,
	}

	return b
}

// format bill (rev)
func (b bill) format() string {
	fs := "Bill's breakdown\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}

	// total
	fs += fmt.Sprintf("%-25v ...$%v \n", "total:", total)

	return fs
}
