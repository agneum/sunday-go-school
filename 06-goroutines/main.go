package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(naturals, squares)
	go printer(squares)

	var input string
	fmt.Scanln(&input)
}

func counter(out chan<- int) {
	for x := 0; ; x++ {
		out <- x
	}
}

func squarer(in <-chan int, out chan<- int) {
	for {
		x := <-in
		out <- x * x
	}
}

func printer(in <-chan int) {
	for {
		fmt.Println(<-in)
	}
}
