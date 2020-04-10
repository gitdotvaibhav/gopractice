package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("->")
		fmt.Println("")

		newstring, _ := reader.ReadString('\n')

		fmt.Println(newstring)

	}

}
