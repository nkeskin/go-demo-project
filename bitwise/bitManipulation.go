package main

import "fmt"

func main() {

	shiftLeft(2, 1)
	shiftRight(16, 2)
}

func shiftLeft(num, shiftCnt int) {
	fmt.Println(num << shiftCnt)
}

func shiftRight(num, shiftCnt int) {
	fmt.Println(num >> shiftCnt)
}
