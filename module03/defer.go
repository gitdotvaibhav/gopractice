package main

import "fmt"

func main() {

	str := "Vaibhav Sanap"

	for _, v := range str {
		defer fmt.Printf("%c", v)
	}

}
