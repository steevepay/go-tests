# NOTES GO

GO DOC: https://golang.org/pkg/

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