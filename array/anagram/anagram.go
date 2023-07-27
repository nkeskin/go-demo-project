package main

import "fmt"

func main() {
	s := "aacc"
	t := "ccac"
	result := isAnagram(s, t)
	fmt.Print(result)
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	myMap := map[rune]int{}
	for _, vS := range s {
		v, ok := myMap[vS]
		if ok {
			v++
			myMap[vS] = v
		} else {
			myMap[vS] = 1
		}
	}
	for _, vT := range t {
		v, ok := myMap[vT]
		v--
		if !ok {
			return false
		} else {
			if v < 0 {
				return false
			}
		}
		myMap[vT] = v
	}
	return true
}
