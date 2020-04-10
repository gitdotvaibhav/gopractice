package main

import "fmt"

func fab(c, quit chan int) {
	x, y := 0, 1
	for {

		select {
		case c <- x:
			x, y = y, x+y

		case <-quit:

			fmt.Printf("Done")
			return

		}
	}

}
func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d\n", <-c)
		}

		quit <- 0

	}()

	fab(c, quit)

}
