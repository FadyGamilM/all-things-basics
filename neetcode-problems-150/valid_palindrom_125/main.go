package main

import (
	"log"
	"strings"
	"unicode"
)

func main() {
	inp := ",-"
	log.Println(isPalindrome(inp))
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	s = strings.ToLower(s)

	l, h := 0, len(s)-1
	for l < h {
		for !(unicode.IsLetter(rune(s[l])) || unicode.IsNumber(rune(s[l]))) && l < len(s)-1 {

			l++
		}

		for !(unicode.IsLetter(rune(s[h])) || unicode.IsNumber(rune(s[h]))) && h > 0 {
			h--
		}

		if l >= h {
			return true
		}

		lowChar := s[l]
		hiChar := s[h]
		lowCharPrint := string(lowChar)
		_ = lowCharPrint
		hiCharPrint := string(hiChar)
		_ = hiCharPrint
		if lowChar != hiChar {
			return false
		} else {
			l++
			h--
		}
	}

	return true
}
