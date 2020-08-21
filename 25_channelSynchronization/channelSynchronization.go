package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("starting work...\n")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)
	// if you removed the <- done line from this program, the program would exit before the worker even started.
	<-done
}
