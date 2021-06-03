package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.Fields("Fields splits the string s around each instance")

	for i, v := range s {
		fmt.Println(i+1, " : ", v)
	}
}
