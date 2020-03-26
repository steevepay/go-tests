# NOTES GO

GO DOC: https://golang.org/pkg/
Examples: https://gobyexample.com/

## commands
```bash
# Compile
$ go build filename.go
# compile and execute
$ go run filenames.go
```

## Package

- Always declare the package name at the top of the GO file.
- Package: a group of files and codes, there is 2 types
  - Package main: will create an binary file
  - Package OtherName: will create a Library/Helper

## import

```go
import "fmt"
import "personnalPackage"
```

## Variable

```go
var card string = "Long version"
var card = "short"
card := "this is also working"
```

Zero values default:
string = ""
int = 0
float = 0
bool = false

## Array

- Array: fix size
- Slice: dynamic array that can grow and shrink. Les données doivent être du même type.

## func

```go
// receiver (d deck)
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// then you can call the function:
cards.print();
```

## Test

les noms de fichiers:
xxx_test.go

Dans le fichier créer des functions type:

```go
package main

import (
	"testing"
)

// Xxxx is the name of the function tested
func TestXxxxx(t *testing.T) {
  // if something is wrong, call:
  t.Errorf("Description")
})
```

## Structs

Examples: https://gobyexample.com/structs

```go
type person struct {
  firstname string
  lastname string
  age int
}

// initialise 1
alex := person{firstname: "Alex", lastname: "fewfew", age: 10}
// initialise 2
var alex person
```

## Pointers

When we pass variables to a function, GO in memory is copying the variables. If we want make changes, we must use a reference with a pointer.
Must use pointer to change the values in a function:
- int, float, string, bool, structs
Must not use and be worry about pointer in a function (it is still copying but it is keeping a reference to the main data):
- slices, maps, channels, pointers, functions
For example, the type slice: An array and a structure that records the length of the slice, the capacity of the slice, and a reference to the underlying array. When we pass a slice,to a function, we are copying the structure, but the reference to the underlying array remains the same.

## Maps

Collection of key value pair (like an object in javascript).
The keys must be all the same type.
The values must be all the same type.
The key type can be different to values type.
Reference Type.

```go
animals:= map[string]int{
  "dog":23,
  "cat":40,
  "cow":22,
}

// Enveler une propriété
remove(animals, "dog");
```

## Interfaces

- Interfaces are not generic types
- Interfaces are implicite, we don't have to say that our custom type satisfies some interface
- Interfaces are not a requirement, sometimes taugh to use.
- We can combine multiple interfaces into a single one.

```go
type bot interface {
	getGreating() string
}
// Now all functions getGreating are heritating from the bot type.
```

## Go Routines

Concurency, we can have multiple threads executing code. If one thred block, another one is picked up and worked on.
Parallelism - Multiple threads executed at the exact same time. Require CPU's

Utiliser les channels pour communiquer entre les routines.
```go
package main
import "fmt"
func main() {
// Create a new channel with make(chan val-type). Channels are typed by the values they convey.
    messages := make(chan string)
// Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
    go func() { messages <- "ping" }()
// The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
    msg := <-messages
    fmt.Println(msg)
}

```