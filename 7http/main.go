package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, resp.Body)

	// buff := make([]byte, 100)
	// for {
	// 	_, errr := resp.Body.Read(buff)

	// 	fmt.Print(string(buff))
	// 	if errr == io.EOF {
	// 		break
	// 	}
	// }

	// fmt.Println(string(body))
}

func (logWriter) Write(bs []byte) (int, error) {
	return 1, nil
}
