package main

import (
	"testing"
)

func TestShiftLeft(t *testing.T) {
	result := ShiftLeft(10, 5)
	expected := 320
	if result != expected {
		t.Errorf("result is : %d , should be : %d", result, expected)
	}
}

func TestShiftRight(t *testing.T) {
	result := ShiftRight(10, 5)
	expected := 0
	if result != expected {
		t.Errorf("result is : %d , should be : %d", result, expected)
	}
}

func TestShiftLeftTableDriven(t *testing.T) {
	var testCases = []struct {
		name     string
		num      int
		shiftCnt int
		result   int
	}{
		{"test1", 1, 2, 4},
		{"test2", 3, 3, 24},
		{"test3", 5, 5, 160},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t2 *testing.T) {
			result := ShiftLeft(tt.num, tt.shiftCnt)
			if result != tt.result {
				t2.Errorf("expected : %d, actual : %d", tt.result, result)
			}
		})
	}
}

func TestShiftRightTableDriven(t *testing.T) {
	var testCases = []struct {
		name     string
		num      int
		shiftCnt int
		result   int
	}{
		{"test1", 100, 4, 6},
		{"test2", 25, 5, 0},
		{"test3", 138, 2, 34},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t2 *testing.T) {
			result := ShiftRight(tt.num, tt.shiftCnt)
			if result != tt.result {
				t2.Errorf("expected : %d, actual : %d", tt.result, result)
			}
		})
	}
}
