package main

import (
	"fmt"
	"time"
)

func pinger(c chan string, pongChannel chan bool) {
	for {
		if !<-pongChannel {
			for i := 1; i <= 2; i++ {
				c <- "ping"
			}
			pongChannel <- true
		}
	}
}

func ponger(c chan string, pongChannel chan bool) {
	for {
		if <-pongChannel {
			c <- "pong"
		}
		pongChannel <- false
	}
}

func printer(c chan string, pongChannel chan bool) {
	pongChannel <- false
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

// Run twice pinger and once ponger
func main() {
	var c chan string = make(chan string, 3)
	var pongChannel chan bool = make(chan bool)

	go pinger(c, pongChannel)
	go ponger(c, pongChannel)
	go printer(c, pongChannel)

	var input string
	fmt.Scanln(&input)
}
