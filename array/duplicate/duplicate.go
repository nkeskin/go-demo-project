package main

import "fmt"

func main() {
	nums := []int{1, 1, 2, 3, 4, 5}
	duplicate := containsDuplicate(nums)
	fmt.Print(duplicate)
}

func containsDuplicate(nums []int) bool {
	m := map[int]int{}
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			return true
		}
		m[v] = v
	}
	return false
}
