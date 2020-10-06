package four

import (
	"fmt"
	"unicode"
)

// exercise 4.5
func deduplicateAdjacent(s []string) []string {
	length := len(s)
	removed := 0
	i := 0
	for i < length-removed-1 {
		if s[i] == s[i+1] {
			copy(s[i:length-removed-1], s[i+1:length-removed])
			removed++
		} else {
			i++
		}
	}
	return s[:length-removed]
}

// exercise 4.6
func squashSpace(s []byte) []byte {
	length := len(s)
	removed := 0
	i := 0
	for i < length-removed-1 {
		if unicode.IsSpace(rune(s[i])) {
			j := 1
			for i+j < length-removed {
				if s[i+j] == s[i] {
					j++
					if i+j == length-removed {
						removed++
					}
				} else {
					if j > 1 {
						copy(s[i:length-removed-j+1], s[i+j-1:length-removed])
						removed += j - 1
					}
					break
				}
			}
			i++
		} else {
			i++
		}
	}
	fmt.Println(s, removed, length)
	return s[:length-removed]
}
