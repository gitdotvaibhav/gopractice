package main

import (
	"fmt"
	"strings"
)

func main() {

	var str strings.Builder

	for i := 0; i < 10; i++ {
		str.WriteString("q")

	}

	fmt.Printf("%s\n", str.String())

}
