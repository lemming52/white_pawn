// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"crypto/sha256" //!+
	"crypto/sha512"
	"fmt"
	"math/bits"
	"os"
)

const (
	MATCH = "-m"
	HASH  = "-h"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Incorrect arguments")
		return
	}
	function := os.Args[1]

	if function == MATCH {
		fmt.Println(matchCount(os.Args[2], os.Args[3]))
	} else if function == HASH {
		hash(os.Args[2], os.Args[3])
	}
}

func matchCount(x, y string) int {
	a := sha256.Sum256([]byte(x))
	b := sha256.Sum256([]byte(y))
	count := 0
	for i, aByte := range a {
		bByte := b[i]
		count += bits.OnesCount8(^(aByte ^ bByte))
	}
	return count
}

func hash(flag, x string) {
	if flag == "256" {
		fmt.Println(sha256.Sum256([]byte(x)))
	} else if flag == "384" {
		fmt.Println(sha512.Sum384([]byte(x)))
	} else if flag == "512" {
		fmt.Println(sha512.Sum512([]byte(x)))
	}
}

//!-
