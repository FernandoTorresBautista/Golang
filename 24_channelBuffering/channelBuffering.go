package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "chan string 1"
	messages <- "chan string 2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
