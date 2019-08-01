package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", " "
	//Simple for loop
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
	s = ""

	//Looping with range and slice
	for _, arg := range os.Args[1:] { //range outputs index and value, to ignore index value we use _ (blank identifier)
		s += sep + arg //creates a new string everytime its concatenated
	}
	fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " ")) //Only one string - memory effficient
}
