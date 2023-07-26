package main

func main() {
	GetBit(24, 3)
}

func ShiftLeft(num, shiftCnt int) int {
	return num << shiftCnt
}

func ShiftRight(num, shiftCnt int) int {
	return num >> shiftCnt
}

func GetBit(num, ith int) bool {
	val := 1 << ith
	if val&num == 0 {
		return false
	}
	return true
}

func SetBit(num, ith int) int {
	val := 1 << ith
	return num | val
}
