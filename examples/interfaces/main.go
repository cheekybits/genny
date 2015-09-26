package main

import "fmt"

//go:generate genny -pkg=main -in=join.go -out=gen-$GOFILE gen "Stringer=MyStr"

type MyStr string

func (s MyStr) String() string {
	return "\"" + string(s) + "\""
}

func main() {
	strs := []MyStr{
		MyStr("hello"),
		MyStr("world"),
	}
	fmt.Printf("%s\n", JoinMyStrs(strs, " "))
}
