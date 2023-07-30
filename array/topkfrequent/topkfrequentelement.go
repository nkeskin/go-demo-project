package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	result := topKFrequent(nums, 2)
	for _, v := range result {
		fmt.Println(v)
	}
}

func topKFrequent(nums []int, k int) []int {
	myMap := map[int]int{}
	var result []int

	for i := 0; i < len(nums); i++ {
		v, _ := myMap[nums[i]]
		v++
		myMap[nums[i]] = v
	}

	for i := 0; i < k; i++ {
		max := 0
		keyFound := 0
		for key, val := range myMap {
			if val >= max {
				max = val
				keyFound = key
			}
		}
		result = append(result, keyFound)
		myMap[keyFound] = 0
	}

	return result
}
