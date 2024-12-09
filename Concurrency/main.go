package main

import (
	"fmt"
	"time"
)

func greet(text string, doneChan chan bool) {
	fmt.Println(text)
	doneChan <- true
}

func slowGreet(text string, doneChan chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println(text)
	doneChan <- true
	close(doneChan) // close channel
}

func main() {
	// dones := make([]chan bool, 4)
	done := make(chan bool)

	// dones[0] = make(chan bool)
	go greet("nice to meet you", done)
	// dones[1] = make(chan bool)
	go greet("how are you", done)
	// dones[2] = make(chan bool)
	go slowGreet("how ... are ... you", done)
	// dones[3] = make(chan bool)
	go greet("I hope you like taking the course", done)

	// for _, done := range dones {
	// 	<-done
	// }

	for _ = range done {
		// fmt.Println(doneChan)
	}

}
