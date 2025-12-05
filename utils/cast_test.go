package utils

import (
	"testing"
)

func TestCast_ToInt_ValidString(t *testing.T) {
	input := "123"
	expected := 123

	actual := ToInt(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCast_ToInt_InvalidString(t *testing.T) {
	input := "abc"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected ToInt to panic for invalid input, but it didn't")
		}
	}()

	ToInt(input)
}

func TestCast_ToIntSlice_ValidStrings(t *testing.T) {
	input := []string{"1", "2", "3", "4", "5"}
	expected := []int{1, 2, 3, 4, 5}

	actual := ToIntSlice(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(actual))
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], actual[i])
		}
	}
}

func TestCast_ToIntSlice_EmptySlice(t *testing.T) {
	input := []string{}
	expected := []int{}

	actual := ToIntSlice(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected empty slice, got length %d", len(actual))
	}
}

func TestCast_ToIntSlice_InvalidString(t *testing.T) {
	input := []string{"1", "abc", "3"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected ToIntSlice to panic for invalid input, but it didn't")
		}
	}()

	ToIntSlice(input)
}
