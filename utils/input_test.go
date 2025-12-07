package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadInput_ValidFile(t *testing.T) {
	// Create a temporary test file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_input.txt")
	content := "line1\nline2\nline3\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadInput(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}
	if result[0] != "line1" || result[1] != "line2" || result[2] != "line3" {
		t.Errorf("Expected [line1, line2, line3], got %v", result)
	}
}

func TestReadInput_EmptyLines(t *testing.T) {
	// Test that internal empty lines are preserved
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_empty.txt")
	content := "line1\n\nline3\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadInput(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}
	if result[0] != "line1" || result[1] != "" || result[2] != "line3" {
		t.Errorf("Expected [line1, , line3], got %v", result)
	}
}

func TestReadInput_WindowsLineEndings(t *testing.T) {
	// Test that Windows line endings are normalized
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_windows.txt")
	content := "line1\r\nline2\r\nline3\r\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadInput(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}
	if result[0] != "line1" || result[1] != "line2" || result[2] != "line3" {
		t.Errorf("Expected [line1, line2, line3], got %v", result)
	}
}

func TestReadInput_NoTrailingNewline(t *testing.T) {
	// Test file without trailing newline
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_no_newline.txt")
	content := "line1\nline2\nline3"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadInput(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}
	if result[0] != "line1" || result[1] != "line2" || result[2] != "line3" {
		t.Errorf("Expected [line1, line2, line3], got %v", result)
	}
}

func TestReadInput_FileNotFound(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when file not found")
		}
	}()

	ReadInput("/nonexistent/path/file.txt")
}

func TestReadRaw_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_raw.txt")
	content := "line1\nline2\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadRaw(testFile)

	if result != content {
		t.Errorf("Expected %q, got %q", content, result)
	}
}

func TestReadLines_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_lines.txt")
	content := "line1\nline2\nline3\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadLines(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(result))
	}
}

func TestReadGrid_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_grid.txt")
	content := "abc\ndef\nghi\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadGrid(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}
	if len(result[0]) != 3 {
		t.Errorf("Expected 3 columns in row 0, got %d", len(result[0]))
	}
	if result[0][0] != 'a' || result[1][1] != 'e' || result[2][2] != 'i' {
		t.Errorf("Grid values incorrect")
	}
}

func TestReadMatrixInt_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_matrix.txt")
	content := "1 2 3\n4 5 6\n7 8 9\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadMatrixInt(testFile, " ")

	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}
	if len(result[0]) != 3 {
		t.Errorf("Expected 3 columns in row 0, got %d", len(result[0]))
	}
	if result[0][0] != 1 || result[1][1] != 5 || result[2][2] != 9 {
		t.Errorf("Matrix values incorrect")
	}
}

func TestReadInts_ValidFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_ints.txt")
	content := "123\n456\n789\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadInts(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 integers, got %d", len(result))
	}
	if result[0] != 123 || result[1] != 456 || result[2] != 789 {
		t.Errorf("Expected [123, 456, 789], got %v", result)
	}
}

func TestReadInts_WithEmptyLines(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test_ints_empty.txt")
	content := "123\n\n456\n\n789\n"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result := ReadInts(testFile)

	if len(result) != 3 {
		t.Errorf("Expected 3 integers (empty lines skipped), got %d", len(result))
	}
	if result[0] != 123 || result[1] != 456 || result[2] != 789 {
		t.Errorf("Expected [123, 456, 789], got %v", result)
	}
}
