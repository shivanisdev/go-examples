package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// We have a string s. You have to find the length of maximum substring with no repeated elements.
	// Example:
	// abadcabc
	// Output = 4 // badc

	str := "abadcabc"

	// to get all substring
	len := len(str)
	sliceSubStr := []string{}
	i := 1
	for i <= len {
		j := i
		for k := 0; k <= i-1; k++ {
			sliceSubStr = append(sliceSubStr, str[k:len+1-j])
			j--
		}
		i++
	}
	fmt.Println("all Substrings =>", sliceSubStr)
	var result string
	for _, ss := range sliceSubStr {
		sCount := map[string]int{}
		flag := true
		for _, v := range ss {
			if _, ok := sCount[string(v)]; ok {
				sCount[string(v)]++
				flag = false
				break
			} else {
				sCount[string(v)] = 1
			}
		}
		if flag {
			result = ss
			break
		}
	}
	fmt.Printf("result if of type %T and value is %v", result, result)
	fmt.Println("-----------")
	fmt.Println("length", utf8.RuneCountInString(result))

}
