package main

import "fmt"

func main() {
	for _, v := range groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) {
		fmt.Println(v)
	}
}

func groupAnagrams(strs []string) [][]string {
	var group [][]string
	firstItem := true
	anagramSet := false
	for i := 0; i < len(strs)-1; i++ {
		var subGroup []string
		for j := i + 1; j < len(strs); j++ {
			if isAnagram(strs[i], strs[j]) && (strs[i] != "#" || strs[j] != "#") {
				if firstItem {
					subGroup = append(subGroup, strs[i])
					firstItem = false
				}
				subGroup = append(subGroup, strs[j])
				strs[j] = "#"
				anagramSet = true
			}
		}
		if anagramSet {
			strs[i] = "#"
			anagramSet = false
		}
		firstItem = true
		if subGroup != nil {
			group = append(group, subGroup)
		}
	}

	for i := 0; i < len(strs); i++ {
		if strs[i] != "#" {
			var subGroup []string
			subGroup = append(subGroup, strs[i])
			group = append(group, subGroup)
		}
	}
	return group
}
