package main

import (
	"fmt"
	"os"
)

/*
Assignment:
Modifying previous echo programs,
print the command in os.Args[0]
print each argument's index and value line by line
*/
func main() {
	s, sep := "", ""
	for idx, arg := range os.Args[0:] {
		sep = " "
		s = arg + sep
		fmt.Println(idx, s)
	}

}
