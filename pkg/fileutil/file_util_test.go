package fileutil

import (
	"errors"
	"io"
	"os"
	"testing"
)

func TestNewFileReader(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testData := []byte("Hello, World!")
	if _, err := tempFile.Write(testData); err != nil {
		t.Fatalf("Write temp file failed: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Close temp file failed: %v", err)
	}

	reader, err := NewFileReader(tempFile.Name())
	if err != nil {
		t.Fatalf("Create FileReader failed: %v", err)
	}
	defer reader.fil.Close()

	data, err := reader.Reader.ReadString('\n')
	if err != nil && err != io.EOF {
		t.Fatalf("ReadString failed: %v", err)
	}

	if data != string(testData) {
		t.Errorf("Read data error, Expected: %q, Real: %q", string(testData)+"\n", data)
	}
}

func TestFileReader_ReadLine(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testData := []byte("Hello, World!\r\nThis is a test.\nAnother line.")
	if _, err := tempFile.Write(testData); err != nil {
		t.Fatalf("Write temp file failed: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Close temp file failed: %v", err)
	}

	reader, err := NewFileReader(tempFile.Name())
	if err != nil {
		t.Fatalf("Create FileReader failed: %v", err)
	}
	defer reader.fil.Close()

	// Test first line
	line, err := reader.ReadLine()
	if err != nil && err != io.EOF {
		t.Fatalf("ReadLine failed: %v", err)
	}
	if line != "Hello, World!" {
		t.Errorf("ReadLine error, Expected: %q, Got: %q", "Hello, World!", line)
	}

	// Test second line with \r\n
	line, err = reader.ReadLine()
	if err != nil && err != io.EOF {
		t.Fatalf("ReadLine failed: %v", err)
	}
	if line != "This is a test." {
		t.Errorf("ReadLine error, Expected: %q, Got: %q", "This is a test.", line)
	}

	// Test third line with \n only
	line, err = reader.ReadLine()
	if err != nil && err != io.EOF {
		t.Fatalf("ReadLine failed: %v", err)
	}
	if line != "Another line." {
		t.Errorf("ReadLine error, Expected: %q, Got: %q", "Another line.", line)
	}

	// Test end of file
	line, err = reader.ReadLine()
	if err != io.EOF {
		t.Errorf("Expected EOF, got: %v", err)
	}
	if line != "" {
		t.Errorf("Expected empty string at EOF, got: %q", line)
	}
}

func TestFileReader_Offset(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testData := []byte("Hello, World!")
	if _, err := tempFile.Write(testData); err != nil {
		t.Fatalf("Write temp file failed: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Close temp file failed: %v", err)
	}

	fr, err := NewFileReader(tempFile.Name())
	if err != nil {
		t.Fatalf("Create FileReader failed: %v", err)
	}
	defer fr.fil.Close()

	// Initial offset should be 0
	if offset := fr.Offset(); offset != 0 {
		t.Errorf("Expected initial offset to be 0, got: %d", offset)
	}

	// Read some data and check the offset
	if _, err := fr.ReadLine(); err != nil && err != io.EOF {
		t.Fatalf("ReadString failed: %v", err)
	}
	if offset := fr.Offset(); offset != int64(len(testData)) {
		t.Errorf("Expected offset to be %d, got: %d", len(testData), offset)
	}
}

func TestFileReader_SeekOffset(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	testData := []byte("Hello, World!\r\nThis is a test.\nAnother line.")
	if _, err := tempFile.Write(testData); err != nil {
		t.Fatalf("Write temp file failed: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		t.Fatalf("Close temp file failed: %v", err)
	}

	fr, err := NewFileReader(tempFile.Name())
	if err != nil {
		t.Fatalf("Create FileReader failed: %v", err)
	}
	defer fr.fil.Close()

	tests := []struct {
		offset       int64
		expect       string
		expectOffset int64
	}{
		{0, "Hello, World!", 15},
		{15, "This is a test.", 31},
		{31, "Another line.", 44},
		{38, " line.", 44},
	}

	for _, test := range tests {
		if err := fr.SeekOffset(test.offset); err != nil {
			t.Errorf("SeekOffset failed at offset %d: %v", test.offset, err)
		}

		line, err := fr.ReadLine()
		if err != nil && err != io.EOF {
			t.Errorf("ReadLine failed after SeekOffset: %v", err)
		}

		if line != test.expect {
			t.Errorf("SeekOffset at %d, Expected: %q, Got: %q", test.offset, test.expect, line)
		}

		if fr.Offset() != test.expectOffset {
			t.Errorf("Offset mismatch after SeekOffset, Expected: %d, Got: %d", test.offset, fr.Offset())
		}
	}
}

func TestFileReader_Close(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	fr, err := NewFileReader(tempFile.Name())
	if err != nil {
		t.Fatalf("Create FileReader failed: %v", err)
	}

	// Close the file reader
	err = fr.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}

	// Attempt to read from the file reader after closing
	_, err = fr.Reader.ReadString('\n')
	if err == nil {
		t.Errorf("Expected error after closing, got nil")
	}

	// Attempt to close the file reader again
	err = fr.Close()
	if err == nil {
		t.Errorf("Expected no error on second close, got: %v", err)
	}
}

func TestIsExist(t *testing.T) {
	// Test for an existing file
	existingFilePath := "test_data/test_1.txt"
	if !IsExist(existingFilePath) {
		t.Errorf("Expected IsExist(%q) to be true", existingFilePath)
	}

	// Test for a non-existing file
	nonExistingFilePath := "/path/to/non/existing/file"
	if IsExist(nonExistingFilePath) {
		t.Errorf("Expected IsExist(%q) to be false", nonExistingFilePath)
	}

	// Test handling of other errors
	// This is a bit tricky since os.Stat doesn't easily simulate other errors.
	// You might need to create a temporary directory and make it unreadable to simulate an error.
	tempDir, err := os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	unreadableFilePath := tempDir + "/unreadable"
	// Make the directory unreadable
	err = os.Chmod(tempDir, 0000)
	if err != nil {
		t.Fatalf("Failed to chmod temp dir: %v", err)
	}

	_, err = os.Stat(unreadableFilePath)
	if !errors.Is(err, os.ErrPermission) {
		t.Fatalf("Expected os.ErrPermission, got %v", err)
	}
	if IsExist(unreadableFilePath) {
		t.Errorf("Expected IsExist(%q) to be false due to permission error", unreadableFilePath)
	}
}

func TestCreateFile(t *testing.T) {
	// Test case 1: File creation success
	tempFilePath := "file.txt"
	success := CreateFile(tempFilePath)
	if !success {
		t.Errorf("Expected file creation to succeed, but it failed")
	}
	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		t.Errorf("File was not created")
	}
	os.Remove(tempFilePath)

	// Test case 2: File creation failure (invalid path)
	tempFilePath = "/invalid/path/that/does/not/exist/file.txt"
	success = CreateFile(tempFilePath)
	if success {
		t.Errorf("Expected file creation to fail, but it succeeded")
	}
}

func TestCreateDir(t *testing.T) {
	// Test creating a new directory
	dirPath := "/tmp/testdir"
	err := CreateDir(dirPath)
	if err != nil {
		t.Fatalf("CreateDir failed: %v", err)
	}
	defer os.RemoveAll(dirPath) // Clean up the created directory

	// Check if the directory exists
	if !IsExist(dirPath) {
		t.Errorf("Directory %s does not exist", dirPath)
	}

	// Test creating a directory that already exists
	err = CreateDir(dirPath)
	if err != nil {
		t.Errorf("CreateDir on existing directory failed: %v", err)
	}

	// Test creating a directory with invalid path
	invalidPath := "/invalid/path"
	err = CreateDir(invalidPath)
	if err == nil {
		t.Errorf("Expected error for invalid path, got nil")
	}
}
