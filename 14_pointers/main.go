package main

import "fmt"

// pass group a type by reference with a pointer
func updateName(n *string) {
	*n = "Mario"
}

func main() {
	n := "mario"

	// store memory location in a pontier
	m := &n
	fmt.Println("memory address contained by m:", m)
	fmt.Println("value at memory address contained by m:", *m)

	// update n with pointer
	updateName(m)
	fmt.Println("name =", n)
}
