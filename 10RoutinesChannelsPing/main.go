package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkURL(c chan string, url string) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url + "might be down!")
	} else {
		fmt.Println(url + " is up!")
	}
	c <- url
}

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, url := range links {
		go checkURL(c, url)
	}

	// SOLUTION 1
	// range similar to <-c
	for l := range c {
		// Fonction literal == Anonymous function in JS
		go func(url string) {
			time.Sleep(time.Second)
			checkURL(c, url)
		}(l)
	}
	// SOLUTION 2
	// for {
	// 	// Blocking call, waiting for the message from the channel
	// 	go checkURL(c, <-c)
	// }
}
