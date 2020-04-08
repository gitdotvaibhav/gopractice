package main

import (
	"sync"
)

type rect struct {
	len   int
	width int
}

func (r rect) area(wg *sync.WaitGroup) int {
	defer wg.Done()

	if r.len == 0 {

		return 0
	}

	if r.width == 0 {

		return 0
	}

	return r.width * r.len

}

func main() {

	var wg sync.WaitGroup

	r1 := rect{len: 2, width: 3}
	r2 := rect{len: 3, width: 10}

	rects := []rect{r1, r2}

	for _, v := range rects {
		wg.Add(1)

		v.area(&wg)

	}

}
