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
	"bytes"
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var buf bytes.Buffer
	
	n := len(s)
	r := n / 3   // 2
	mod := n % 3 // 2
	if mod == 0 {
		mod = 3
		r--
	}	

	buf.WriteString(s[:mod])
	// 12 123 123 n=8
	// 12,
	
	for i := 0; i < r; i++ {
		buf.WriteString(",")
		buf.WriteString(s[mod+i*3:mod+i*3+3])
	}
	return buf.String()
}

//!-
