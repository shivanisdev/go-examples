package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := os.Args[1]
	fmt.Println(strings.Repeat("!", len(s)) + s + strings.Repeat("!", len(s)))
}
