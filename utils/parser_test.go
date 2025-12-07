package utils

import "testing"

func TestParse_ExtractInts_ValidLine(t *testing.T) {
	line := "123 456 789"
	res := ExtractInts(line)
	if len(res) != 3 {
		t.Errorf("Expected 3 integers, got %d", len(res))
	}
	if res[0] != 123 || res[1] != 456 || res[2] != 789 {
		t.Errorf("Expected [123, 456, 789], got %v", res)
	}
}

func TestParse_ExtractInts_EdgeCases(t *testing.T) {
	if ExtractInts("abc def 123")[0] != 123 {
		t.Errorf("Expected 123, got %d", ExtractInts("abc def 123")[0])
	}

	res := ExtractInts("123,345.567")
	if res[0] != 123 {
		t.Errorf("Expected 123, got %d", res)
	}
	if res[1] != 345 {
		t.Errorf("Expected 345, got %d", res)
	}
	if res[2] != 567 {
		t.Errorf("Expected 567, got %d", res)
	}
}

func TestParse_ParsePattern_HappyCase(t *testing.T) {
	line := "abc123def456ghi789"
	pattern := "([a-z]+)([0-9]+)([a-z]+)"
	res := ParsePattern(line, pattern)
	if len(res) != 3 {
		t.Errorf("Expected 3 matches, got %d", len(res))
	}
	if res[0] != "abc" || res[1] != "123" || res[2] != "def" {
		t.Errorf("Expected [abc, 123, def], got %v", res)
	}
}

func TestParse_ParsePattern_EdgeCases(t *testing.T) {
	line := "abc123def456ghi789"
	pattern := "([a-z]+)([0-9]+)([a-z]+)"
	res := ParsePattern(line, pattern)
	if len(res) != 3 {
		t.Errorf("Expected 3 matches, got %d", len(res))
	}
	if res[0] != "abc" || res[1] != "123" || res[2] != "def" {
		t.Errorf("Expected [abc, 123, def], got %v", res)
	}
}

func TestParse_ParsePattern_NoMatch(t *testing.T) {
	// Test when pattern doesn't match - should return nil
	line := "abc123def"
	pattern := "([0-9]+)xyz([a-z]+)"
	res := ParsePattern(line, pattern)
	if res != nil {
		t.Errorf("Expected nil when no match, got %v", res)
	}
}

func TestParse_ParsePattern_NoCapturingGroups(t *testing.T) {
	// Test when pattern matches but has no capturing groups - should return nil
	line := "abc123def"
	pattern := "abc[0-9]+def"
	res := ParsePattern(line, pattern)
	if res != nil {
		t.Errorf("Expected nil when no capturing groups, got %v", res)
	}
}
