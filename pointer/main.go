package main

import "fmt"

func main() {
	b := "aaaa"

	work(&b)

	fmt.Print(b)
}

func work(s *string) {
	ss := "bbbbb"
	s = &ss
} 