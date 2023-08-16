package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 0}
	fmt.Println(threeSumV2(nums))
}

func threeSumV2(nums []int) [][]int {
	slices.Sort(nums)
	numMap := map[int]int{}
	for i := range nums {
		if nums[i] >= 0 {
			numMap[nums[i]] = i
		}
	}
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 {
			if nums[i] == nums[i-1] {
				continue
			}
		}
		if nums[i] > 0 {
			return result
		}
		target := -nums[i]
		for j := i + 1; j < len(nums)-1; j++ {
			val, ok := numMap[target-nums[j]]
			if ok && val != j {
				subResult := []int{nums[i], nums[val], nums[j]}
				if !containsSubResult2(result, subResult) {
					result = append(result, subResult)
				}
			}
		}
	}
	return result
}

func containsSubResult2(result [][]int, subResult []int) bool {
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
