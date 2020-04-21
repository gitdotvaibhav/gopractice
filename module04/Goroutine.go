package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func doSomething(killChannel chan int) {

	var loopCount int64

	for {

		loopCount = loopCount + 1
		fmt.Printf("Loop Count %d", loopCount)
		time.Sleep(10 * time.Second)

		select {

		case <-killChannel:
			fmt.Println("Exiting the goroutine")
			return

		default:
			fmt.Println("Continue")
		}

	}

}

func main() {

	killChannel := make(chan int, 10)
	breakChannel := make(chan int, 1)

	prev := 0
	iterator := 0
	for {

		file, err := os.Open("count.txt")
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		scanner := bufio.NewScanner(file)

		var currentCount int
		go func() {

			currentCount, err = strconv.Atoi(scanner.Text())
			if err == nil {
				breakChannel <- 1
			}

			select {
			case <-breakChannel:
				return
			default:
			}
		}()
		file.Close()
		fmt.Printf("Number of go routine should be %d\n", currentCount)

		if err != nil {
			fmt.Println("invalide input")
			currentCount = prev
		}

		if prev > currentCount {
			for iterator = 0; iterator < prev-currentCount; iterator++ {
				killChannel <- 1
			}
		}

		if prev < currentCount {
			for iterator = 0; iterator < currentCount-prev; iterator++ {
				fmt.Printf("Creating go routine \n")
				go doSomething(killChannel)
			}
		}

		prev = currentCount

		time.Sleep(2 * time.Second)
	}

}
