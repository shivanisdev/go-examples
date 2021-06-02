package main

import (
	"fmt"
	"path"
)

func main() {
	// Base

	fmt.Println(path.Base("/a/b"))
	fmt.Println(path.Base("https://golang.org/pkg/path/"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))

	fmt.Print("\n")
	//clean

	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}

	fmt.Print("\n")
	//Dir

	fmt.Println(path.Dir("/a/b/c"))
	fmt.Println(path.Dir("a/b/c"))
	fmt.Println(path.Dir("/a/"))
	fmt.Println(path.Dir("a/"))
	fmt.Println(path.Dir("/"))
	fmt.Println(path.Dir(""))

	fmt.Print("\n")
	//Ext

	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Ext("/"))
	fmt.Println(path.Ext(""))

	fmt.Print("\n")
	//IsAns

	fmt.Println(path.IsAbs("/dev/null"))
	fmt.Print(path.IsAbs("https://github.com/golang/go/blob/master/src/path/path.go"))

	fmt.Print("\n\n")
	//Join

	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))

	fmt.Println(path.Join("a/b", "../../../xyz"))

	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))

	fmt.Print("\n\n")
	//Spit

	split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("static/myfile.css")
	split("myfile.css")
	split("")
	split("https://golang.org/pkg/path/")

}
