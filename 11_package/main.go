package main

import "fmt"

var score = 99.5

func main() {
	sayHello("Mario")

	for _, p := range points {
		fmt.Println("point", p)
	}

	showScore()
}
