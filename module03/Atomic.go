package main

import (
	"sync"
	"sync/atomic"

	"runtime"
)

var (
	counter int64
	wg      *sync.WaitGroup
)

func main() {

	go increamentCounter()
	go increamentCounter()
}

func increamentCounter() {

	atomic.AddInt64(&counter, 4)
	runtime.Gosched()
}
