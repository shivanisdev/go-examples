package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	//fmt.Println(join("abc", "123", "xyz", "abc", "123", "xyz", "abc", "123", "xyz", "abc", "123", "xyz", "abc", "123", "xyz"))
	fmt.Println(join2("abc", "123", "xyz", "abc", "123", "xyz", "abc", "123", "xyz", "abc", "123", "xyz", "abc", "123", "xyz"))
	fmt.Println(time.Since(start).Nanoseconds())
}

func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

func join2(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
