package main

import "fmt"

func sumofTwoLista(list1 []int, list2 []int) int {

	total := 0
	for val := range append(list1, list2...) {
		total = total + val
	}

	return total
}

func main() {

	list1 := []int{1, 2, 3, 4, 5}
	list2 := []int{6, 7, 2, 4, 2, 4, 14, 1}
	fmt.Printf("Combined Sum of list : %v \n", sumofTwoLista(list1, list2))

}
