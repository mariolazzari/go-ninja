package main

import (
	"fmt"
	"math"
)

func greeting(name string) {
	fmt.Printf("Hello %v\n", name)
}

func sayBye(name string) {
	fmt.Printf("Bye %v\n", name)
}

func cycleNames(names []string, f func(string)) {
	for _, v := range names {
		f(v)
	}
}

// return value
func circleArea(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func main() {
	greeting("Mario")
	sayBye("Mariarosa")

	names := []string{"Mario", "Mariarosa", "Maria"}
	cycleNames(names, greeting)

	area1 := circleArea(5)
	area2 := circleArea(10)
	fmt.Printf("area1 = %0.3f, area2 = %.3f\n", area1, area2)
}
