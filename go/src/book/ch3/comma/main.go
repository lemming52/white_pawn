package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
		fmt.Printf("  %s\n", bufferedComma(os.Args[i]))
	}
}

//!+
func comma(s string) string {
	substrings := strings.Split(s, ".")
	switch len(substrings) {
	case 1:
		if !unicode.IsDigit(rune(s[0])) {
			return string(s[0]) + comma(s[1:])
		}
		n := len(s)
		if n <= 3 {
			return s
		}
		return comma(s[:n-3]) + "," + s[n-3:]
	case 2:
		return comma(substrings[0]) + "." + substrings[1]
	default:
		fmt.Print("To many '.', aborting.")
		return ""
	}
}

func bufferedComma(s string) string {
	var buf bytes.Buffer
	length := len(s)
	for i := 0; i < length; i++ {
		if (length-i)%3 == 0 {
			buf.WriteByte(',')
		}
		buf.WriteByte(s[i])
	}
	return buf.String()
}

//!-
