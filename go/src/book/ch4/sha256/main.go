// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"fmt"
	"math/bits"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	for i, n := range c1 {
		fmt.Printf("%08b %08b\n", n, c2[i]) // prints 00000000 11111101
	}
	fmt.Println(matchCount(c1, c2))
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

func matchCount(a, b [32]byte) int {
	count := 0
	for i, aByte := range a {
		bByte := b[i]
		count += bits.OnesCount8(^(aByte ^ bByte))
	}
	return count
}

//!-
