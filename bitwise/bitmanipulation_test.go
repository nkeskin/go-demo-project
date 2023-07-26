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
