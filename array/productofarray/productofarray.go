package main

import "fmt"

func main() {
	result1 := productExceptSelfWithDivision([]int{1, 2, 3, 4})
	fmt.Println(result1)
	result2 := productExceptSelfWithoutDivision([]int{1, 2, 3, 4})
	fmt.Println(result2)
}

func productExceptSelfWithoutDivision(nums []int) []int {
	prefix := 1
	postfix := 1
	arrLen := len(nums)
	prefixArr := make([]int, arrLen)
	postfixArr := make([]int, arrLen)
	for i := 0; i < arrLen; i++ {
		prefixArr[i] = prefix
		prefix *= nums[i]
	}
	for i := arrLen - 1; i >= 0; i-- {
		postfixArr[i] = postfix
		postfix *= nums[i]
	}
	for i := 0; i < arrLen; i++ {
		prefixArr[i] = prefixArr[i] * postfixArr[i]
	}
	return prefixArr
}

func productExceptSelfWithDivision(nums []int) []int {
	arrLen := len(nums)
	result := make([]int, arrLen)
	mul := 1
	zeroCount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroCount++
		} else {
			mul *= nums[i]
		}
	}

	for i := 0; i < len(nums); i++ {
		if zeroCount == 0 {
			result[i] = mul / nums[i]
		} else {
			if zeroCount == 1 {
				if nums[i] == 0 {
					result[i] = mul
				} else {
					result[i] = 0
				}
			} else {
				result[i] = 0
			}
		}
	}

	return result
}
