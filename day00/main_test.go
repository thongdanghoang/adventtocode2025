package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"1000",
		"2000",
		"3000",
	}
	expected := 6000
	if result := Part1(input); result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
