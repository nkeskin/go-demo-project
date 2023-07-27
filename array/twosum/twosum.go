package main

import "fmt"

func main() {
	nums := []int{2, 4, 19, 21}
	target := 23
	result := twoSum(nums, target)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func twoSum(nums []int, target int) []int {
	myMap := map[int]int{}
	var indices []int
	for i, v := range nums {
		myMap[v] = i
	}
	for i, v := range nums {
		vMap, ok := myMap[target-v]
		if ok {
			if vMap != i {
				indices = append(indices, i)
				indices = append(indices, vMap)
				return indices
			}
		}
	}
	return indices
}
