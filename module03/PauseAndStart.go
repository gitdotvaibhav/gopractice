package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {

	time.Sleep(250 * time.Millisecond)

	i++

	fmt.Printf("Work done %d", i)

}

func goroutine(command <-chan string, wg *sync.WaitGroup) {

	defer wg.Done()
	var status = "Play"

	for {

		select {
		case cmd <- command:
			fmt.Printf("Command %s", cmd)

		}
	}

}
