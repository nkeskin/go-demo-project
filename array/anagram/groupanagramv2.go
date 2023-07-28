package main

import "fmt"

// TODO add an implementation with lower time complexity
func main() {
	for _, v := range groupAnagramsV2([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) {
		fmt.Println(v)
	}
}

func groupAnagramsV2(strs []string) [][]string {
	var group [][]string
	for _, v := range strs {
		for _, z := range v {
			fmt.Println(z - 'a')
		}
		fmt.Println()
	}
	return group
}
