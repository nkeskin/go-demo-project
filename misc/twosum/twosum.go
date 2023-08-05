package main

import "fmt"

func main() {
	input := []int{2, 7, 11, 15}
	result := twoSum(input, 9)
	fmt.Println(result)
}

// must contain constant space
func twoSum(numbers []int, target int) []int {
	var solution []int
	for i := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				solution = append(solution, i+1, j+1)
			}
		}
	}
	return solution
}
