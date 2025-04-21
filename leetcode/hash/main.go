package main

import "fmt"

func hash(s string) int {
	var h int
	for i := 0; i < len(s); i++ {
		h = 31*h + int(s[i])
	}
	return h
}

func main() {
	s := "a"
	fmt.Println(hash(s))
}
