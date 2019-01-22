package main

import (
	"fmt"
	"os"
)

func main() {
	res := anagram(os.Args[1], os.Args[2])
	if res {
		fmt.Println("Words are anagrams")
	} else {
		fmt.Println("Words are not anagrams")
	}
}

func anagram(a, b string) bool {
	lengtha := len(a)
	if len(b) != lengtha {
		return false
	}
	aMap := make(map[byte]int)
	for i := 0; i < lengtha; i++ {
		_, ok := aMap[a[i]]
		if ok {
			aMap[a[i]]++
		} else {
			aMap[a[i]] = 1
		}
	}
	for i := 0; i < lengtha; i++ {
		aMap[b[i]]--
	}
	for _, v := range aMap {
		if v != 0 {
			return false
		}
	}
	return true
}
