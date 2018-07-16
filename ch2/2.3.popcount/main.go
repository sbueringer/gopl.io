// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main(){
	var x uint64 = 14112321312315463123

	fmt.Printf("Value: %b\n", x)

	res1 := PopCount(x)
	res2 := PopCountPerBit(x)

	fmt.Printf("Count bits table: %d\n", res1)
	fmt.Printf("Count bits per bit: %d\n", res2)



	fmt.Printf("Count bits per bit: %08b\n", 156)
	fmt.Printf("Count bits per bit: %08b\n", 155)
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var sum byte
	for i := uint(0); i<8; i++ {
		fmt.Printf("%08b %v\n", byte(x>>(i*8)), pc[byte(x>>(i*8))])
		sum += pc[byte(x>>(i*8))]
	}
	return int(sum)
}

func PopCountPerBit(x uint64) int {
	var sum byte
	for i := uint(0); i<64; i++ {
		sum += byte(x>>(i)) & 1
	}
	return int(sum)
}

//!-
