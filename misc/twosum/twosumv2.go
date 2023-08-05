package main

import "fmt"

func main() {
	input := []int{2, 7, 11, 15}
	result := twoSumV2(input, 9)
	fmt.Println(result)
}

func twoSumV2(numbers []int, target int) []int {
	var solution []int
	numMap := map[int]int{}
	for i := range numbers {
		numMap[numbers[i]] = i
	}
	for i := range numbers {
		val, ok := numMap[target-numbers[i]]
		if ok {
			solution = append(solution, i+1, val+1)
			return solution
		}
	}
	return solution
}
