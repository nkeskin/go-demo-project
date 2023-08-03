package main

import "fmt"

// TODO Not working -- fix implementation
func main() {
	input := []int{1, 0, -1}
	fmt.Println(longestConsecutive(input))
}

func longestConsecutive(nums []int) int {
	numMap := map[int]int{}
	for i := range nums {
		numMap[nums[i]] = nums[i]
	}
	max := 0
	temp := 1
	init := numMap[nums[0]] + 1
	for i := range nums {
		_, ok := numMap[init]
		if ok {
			if i > 0 {
				i--
			}
			init++
			temp++
		} else {
			init = nums[i+1] + 1
			temp = 1
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
