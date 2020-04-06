package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid use ")
		os.Exit(-1)
	}
}

func main() {

	r, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println("Error")
		return
	}

	io.Copy(os.Stdout, r.Body)

	fmt.Println(err)

}
