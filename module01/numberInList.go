package main

import (
	"fmt"
)

func main() {

	list := [5]int{1, 2, 3, 4, 6}
	fmt.Println(numInList(list, 5))
}

func numInList(list [5]int, num int) bool {

	for val := range list {

		if val == num {
			return true
		}
	}

	return false

}
