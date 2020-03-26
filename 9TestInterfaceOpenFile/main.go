package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		log.Fatal("Error: missing file input")
		os.Exit(1)
	}
	filename := args[0]
	fmt.Println(filename)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error:" + err.Error())
	}
	io.Copy(os.Stdout, file)
}
