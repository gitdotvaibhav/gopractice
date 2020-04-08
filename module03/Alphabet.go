package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	runtime.GOMAXPROCS(2)

	var wg *sync.WaitGroup

	wg.Add(2)

	fmt.Println("Started")

	go func() {

		defer wg.Done()

		for 

	}()

}
