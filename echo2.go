package main

import (
	"fmt"
	"os"
)

// func main() {
func echo2() {
	s, sep := "", ""
	// for idx, arg := range os.Args[1:] {
	for _, arg := range os.Args[1:] {
		// fmt.Println(idx)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
