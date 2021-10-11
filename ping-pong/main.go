package main

import (
	"fmt"
)

func main() {
	fmt.Println("Ping Ping")
	table := make(chan int)

	go player("even", table)

	go player("odd", table)

	table <- 0
	//time.Sleep(1 * time.Second)
	<-table // blocks
}

func player(name string, ch chan int) {
	for {
		table := <-ch
		table++
		fmt.Println(name)
		fmt.Println(table)
		ch <- table
		if table == 10 {
			break
		}
	}
}
