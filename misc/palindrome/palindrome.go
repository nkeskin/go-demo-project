package main

import (
	"fmt"
	"strings"
)

func main() {
	input := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(input))
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	charMap := map[int]string{}
	charMap[0] = ""
	for _, c := range s {
		checkValidChar(c, charMap)
	}

	s = charMap[0]
	charMap = nil
	length := len(s)

	for i := 0; i < length/2; i++ {
		if s[i] != s[length-i-1] {
			return false
		}
	}

	return true
}

func checkValidChar(c rune, charMap map[int]string) {
	if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
		v := charMap[0]
		v = v + string(c)
		charMap[0] = v
	}
}
