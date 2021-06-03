package main

import (
	"fmt"
	"os"
	"strings"
)

const s = "lazy cat run again and run again and again"

func main() {
	ss := strings.Fields(s)

	w := os.Args[1:]

query:
	for _, v := range w {
		for i, v1 := range ss {
			if strings.EqualFold(v, v1) {
				fmt.Printf("#%-2d: %q\n", i+1, v1)
				//break query
				//break
				continue query
			}
		}
	}
}
