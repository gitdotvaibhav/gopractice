package main

import "fmt"

type t int

func isClosed(ch <-chan t) closed bool {

	defer func() {
		if recover() != nil {
			// The return result can be altered
			// in a defer function call.
			justClosed = false
		}
	}()

	select {

	case <-ch:
		return true
	default:

	}

	close(ch)

	return false
}

func main() {

	ch := make(chan t)

	fmt.Printf("channel is %v\n", isClosed(ch))

	close(ch)

	fmt.Printf("channel is %v\n", isClosed(ch))

}
