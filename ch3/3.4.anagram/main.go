// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
)

func main() {
	var charMap map[byte]int = make(map[byte]int)

	string1 := os.Args[1]
	string2 := os.Args[2]

	if len(string1) != len(string2) {
		fmt.Println("false1")
		return
	}

	for i:=0; i<len(string1); i++ {
		charMap[string1[i]]++
		charMap[string2[i]]--
	}

	for _,v := range charMap {
		if v != 0 {
			fmt.Println("false2")
			break
		}
	}
	fmt.Println("true")
}

