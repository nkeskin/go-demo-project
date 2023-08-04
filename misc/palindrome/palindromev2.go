package main

import (
	"fmt"
	"strings"
)

func main() {
	input := ".,"
	fmt.Println(isPalindrome2(input))
}

func isPalindrome2(s string) bool {
	s = strings.ToLower(s)

	length := len(s)
	l := 0
	r := length - 1
	for l < r {
		for !checkValidChar(s[l]) && l < length {
			l++
			if l >= length {
				return true
			}
		}
		for !checkValidChar(s[r]) && r > 0 {
			r--
			if r < 0 {
				return true
			}
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func checkValidChar(c uint8) bool {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		return true
	}
	return false
}
