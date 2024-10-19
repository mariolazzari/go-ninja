package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// strings package
	greeting := "Ciao a tutti!"

	fmt.Println(strings.Contains(greeting, "Ciao"))
	fmt.Println(strings.ReplaceAll(greeting, "Ciao", "Buongiorno"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println((strings.Index(greeting, "a")))
	fmt.Println(strings.Split(greeting, " "))

	// original value is unchanged
	fmt.Println((greeting))

	// sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	// sort integers
	sort.Ints(ages)
	fmt.Println((ages))

	// find index
	EL := 30
	index := sort.SearchInts(ages, EL)
	fmt.Printf("%v index is %v\n", EL, index)

	names := []string{"Mario", "Mariarosa", "Roberto", "Maria"}

	// sort strings
	sort.Strings(names)
	fmt.Println((names))

	// search string
	fmt.Println((sort.SearchStrings(names, "Mario")))
}
