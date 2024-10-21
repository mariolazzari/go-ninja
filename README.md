# Go language by NetNinja

## Introduction & setup

### What is Go?

- Fast, statically typed, compiled language
- General purpose
- Built in testing support
- Object oriented (in its own way)

### Installation

Go [site](https://go.dev/)

Check installation:

```sh
go
```

Install Go extension for VS Code

## Fist program

Every go file must be included in a package.
File included in package called _main_ will be compiled in an executable file.
For libraries, use package names that describe their use.

### Functions

For function declarations, use _func_ keyword.
Function _main_ is the entry point of the compiled program.

### Imports

For importing files use _import_ keyword.
Multiple imports will be put between brackets.
Unused imports will cause errors.

### Go standard library

Go [standard library](https://pkg.go.dev/std) includes lots of different functionalities: _fmt_ is for formatting strings.
All methods of a package start with capital letter: only method with capital letter can be exported from a package.

```go
package main

import "fmt"

func main() {
	fmt.Println(("Hello Mario!"))
}
```

### Compile program

```sh
go run main.go
```

## Variables

Unused variable will cause error (and its import too).
Type can be inferred by variable initialization.

There are several ays to declare a varaible :

- _var_ keyword
- _:=_ operator (in function body only)

### Types

Types [reference](https://go.dev/ref/spec#Types)

### Strings

Strings are created within dounle quotes (single ones not allowed).
A not initialized string will be defaulted to an empty string ("").

String [type](https://go.dev/ref/spec#String_types)

```go
// full declaratin
var nameOne string = "mario"
// type inferred
var nameTwo = "luigi"
// no initialization
var nameThree string

fmt.Println(nameOne, nameTwo, nameThree)
```

### Integers

Integer [variations](https://go.dev/ref/spec#Numeric_types)

```go
var ageOne int = 20
var ageTwo = 30
ageThree := 40

fmt.Println(ageOne, ageTwo, ageThree)
```

### Floats

```go
var scoreOne float32 = 25.98
var scoreTwo float64 = 1965385877.5
var scoreThree = 1.5 // inferred as float64
```

## Printing & Formatting Strings

Package [_fmt_](https://pkg.go.dev/fmt) implements formatted I/O with functions.

- _Print_: outputs a string without ending new line
- _Println_: outputs a string with ending new line
- _Printf_: outputs strings with format identifiers:
  - %v: default
  - %q: puts quotes around string variables
  - %T: prints variable type
  - %f: prints float variables
    - %0.nf: prints n decimals
- _Sprintf_: saves printf output to a variable

```go
// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	// Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name)

	// Printf (formatted string), %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name) // %v = value in default format
	fmt.Printf("my age is %q and my name is %q \n", age, name) // %q = quotes
	fmt.Printf("age is of type %T \n", age)                    // %T is the type
	fmt.Printf("you scored %f points! \n", 225.55)             // %f = float format
	fmt.Printf("you scored %0.1f points! \n", 225.55)          // %0.2f = float with 2 decimal points

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)
```

## Arrays & Slices

### Arrays

Arrays are fixed length: array size must be specified in declaration.
If _no size_ has been specified in declaration, a _slice_ will be created instead of an array.xs
_len_ built-in method returns the length of the array.

```go
// var ages [3]int = [3]int{20, 25, 30}
var ages = [3]int{20, 25, 30}

names := [4]string{"yoshi", "mario", "peach", "bowser"}
names[1] = "luigi"

fmt.Println(ages, len(ages))
fmt.Println(names, len(names))
```

### Slices

Slices can change their length and they do not need size in declaration.
_append_ method appends a new element to slice and returns a new slice.
Slice range _[from:to_excluded]_ selects elements from specified element to the first excluded.
Ranges can be appended.

```go
var scores = []int{100, 50, 60}
scores[2] = 25
scores = append(scores, 85)

fmt.Println(scores, len(scores))

// slice ranges
rangeOne := names[1:4] // doesn't include pos 4 element
rangeTwo := names[2:]  //includes the last element
rangeThree := names[:3]
```

## The Standard Library

[docs](https://pkg.go.dev/std)

### strings

[docs](https://pkg.go.dev/strings)

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	greeting := "Ciao a tutti!"

	fmt.Println(strings.Contains(greeting, "Ciao"))
    fmt.Println(strings.ReplaceAll("Ciao", "Buongiorno))
}
```

### sort

Sort methods alter original slice

[docs](https://pkg.go.dev/sort)

```go
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
```

## Loops

### for

```go
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
```

## Booleans & Conditionals

```go
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
}
```

## Functions

```go

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
```

## Multiple Return Values

```go
package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])

	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, ln1 := getInitials("mario lazzari")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("mariarosa")
	fmt.Println(fn2, ln2)
}
```

## Package scope

```go
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
```

```go
package main

import "fmt"

var points = []int{10, 20, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

func showScore() {
	fmt.Println("Score:", score)
}
```

```sh
go run 11_package/main.go 11_package/greetings.go
```

## Maps

Key / value pairs

```go
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
```

## Pass by value

Go makes a copy of values passed to a function from Group A types.
Group B types are passed by reference.

### Group types

| Group A  | Group B   |
| -------- | --------- |
| Strings  | Slices    |
| Ints     | Maps      |
| Floats   | Functions |
| Booleans |           |
| Arrays   |           |
| Structs  |           |

```go
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
```

## Pointers

| n     | m     |
| ----- | ----- |
| 0x001 | 0x002 |
| mario | 0x001 |

```go
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
```

## Structs

Custom object structure

### Desfinition

```go
package main

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// add new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}
```

### Creation

```go
package main

import "fmt"

func main() {
	myBill := newBill("Mario's bill")
	fmt.Println((myBill))
}
```

## Receiver functions

Function associated to a struct.

```go
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
```

## Receiver functions with pointers

```go
// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
```

## User input

Package [bufio](https://pkg.go.dev/bufio) implements buffered I/O.
Package [os](https://pkg.go.dev/os) provides a platform-independent interface to operating system.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)
	b := newBill(name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose option (a - add an item, s - save bill, t - add tip)", reader)
	fmt.Println(opt, b)
}

func main() {
	myBill := createBill()
	promptOptions(myBill)
}
```

## Switch

```go
switch opt {
	case "a":
		fmt.Println("You chose 'a'")

	case "s":
		fmt.Println("You chose 's'")

	case "t":
		fmt.Println("You chose 't'")

	default:
		fmt.Println("Invalid option...")
		promptOptions(b)
	}
```



