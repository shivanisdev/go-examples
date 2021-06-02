package main

import "fmt"

func main() {
	ch := make(chan int)

	go sender(ch)

	receiver(ch)

	fmt.Println("function ends here")

}

func sender(c chan<- int) {
	c <- 42
}

func receiver(c <-chan int) {
	fmt.Println(<-c)
}
