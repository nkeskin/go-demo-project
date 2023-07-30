package main

import (
	"fmt"
)

// TODO add an implementation with lower time complexity
func main() {
	for _, v := range groupAnagramsV2([]string{"c", "c"}) {
		fmt.Println(v)
	}
}

func groupAnagramsV2(strs []string) [][]string {
	var group [][]string
	charMap := map[[26]int][]string{}
	for i, v := range strs {
		chars := [26]int{}
		for _, c := range v {
			val := chars[c-'a']
			val++
			chars[c-'a'] = val
		}
		val, ok := charMap[chars]
		if ok {
			val = append(val, v)
		} else {
			val = append(val, strs[i])
		}
		charMap[chars] = val
		//fmt.Println(charMap[v])
	}
	for _, val := range charMap {
		group = append(group, val)
	}
	return group
}
