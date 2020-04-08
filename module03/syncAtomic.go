package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var (
	flag int64
)

func main() {

	wg sync.WaitGroup

	go flagCheck()
	go flagCheck()

	time.Sleep(6000)

	atomic.StoreInt64(&flag, 1)

}

func flagCheck() {

	for {

		if atomic.LoadInt64(&flag) == 1 {

			break
		}

		time.Sleep(6000)

	}

}
