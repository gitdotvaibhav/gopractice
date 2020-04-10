package main

import (
	"os"
	"time"
)

//Runner runs set of task within amount of time or shut down an operating system
type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timout    <-chan time.Time
}
