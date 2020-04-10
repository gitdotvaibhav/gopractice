package main

import "fmt"

func reverse(arr []int) {

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {

		arr[i], arr[j] = arr[j], arr[i]
	}

}

func main() {

	arr := []int{1, 2, 3, 1, 3, 1, 3, 1, 3, 5, 2, 411, 2334}

	reverse(arr)

	fmt.Println(arr)

	a := map[string]bool{"abc": true, "b": false}

	b := make(map[string]bool)

	for key, val := range a {
		b[key] = val
	}

	fmt.Println(b)

}
