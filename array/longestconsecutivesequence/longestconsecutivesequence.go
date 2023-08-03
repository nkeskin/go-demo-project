package main

import "fmt"

func main() {
	input := []int{9, 1, -3, 2, 4, 8, 3, -1, 6, -2, -4, 7}
	fmt.Println(longestConsecutive2(input))
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

func longestConsecutive2(nums []int) int {
	numMap := map[int]int{}
	max := 0

	for i := range nums {
		numMap[nums[i]] = nums[i]
	}
	for i := range nums {
		temp := 1
		_, ok := numMap[nums[i]-1]
		if !ok {
			seq := 1
			_, ok2 := numMap[nums[i]+seq]
			for ok2 {
				fmt.Println(nums[i] + seq)
				seq++
				temp++
				_, ok2 = numMap[nums[i]+seq]
			}
		}
		if temp > max {
			max = temp
		}
	}
	return max
}
