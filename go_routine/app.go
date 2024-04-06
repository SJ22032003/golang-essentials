package main

import (
	"fmt"
	"time"
)

func main() {

	done := make(chan bool)

	go greet("world1", done)
	go greet("world2", done)
	go slowGreet("goroutine", done)

	for range done {
		fmt.Println("done")
	}

}

func greet(phrase string, doneChan chan bool) {
	fmt.Println("hello", phrase)
	doneChan <- true
}
func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("hello", phrase)
	done <- true
	close(done)
}
