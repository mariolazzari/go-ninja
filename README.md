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
File included in package called *main* will be compiled in an executable file.
For libraries, use package names that describe their use.

### Functions

For function declarations, use *func* keyword.
Function *main* is the entry point of the compiled program.

### Imports

For importing files use *import* keyword.

### Go standard library

Go standard library includes lots of different functionalities: *fmt* is for formatting strings.
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

