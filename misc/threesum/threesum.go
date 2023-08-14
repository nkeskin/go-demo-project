package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					subResult := []int{nums[i], nums[j], nums[k]}
					if !containsSubResult(result, subResult) {
						result = append(result, subResult)
					}
				}
			}
		}
	}
	return result
}

func containsSubResult(result [][]int, subResult []int) bool {
	for _, v := range result {
		count := 0
		for i := range v {
			if v[i] == subResult[i] {
				count++
			}
		}
		if count == 3 {
			return true
		}
	}
	return false
}
