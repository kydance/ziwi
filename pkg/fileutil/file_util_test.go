package fileutil

import (
	"archive/zip"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
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

func TestCopyFile(t *testing.T) {
	src := "src.txt"
	dst := "dst.txt"

	// Create a source file with test data
	testData := []byte("Hello, World!")
	if err := os.WriteFile(src, testData, 0644); err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	defer os.Remove(src)

	// Ensure the destination file does not exist before the test
	if IsExist(dst) {
		t.Fatalf("Destination file should not exist before the test")
	}

	// Perform the copy operation
	if err := CopyFile(dst, src); err != nil {
		t.Fatalf("CopyFile failed: %v", err)
	}

	// Verify the destination file has been created and contains the correct data
	dstData, err := os.ReadFile(dst)
	if err != nil {
		t.Fatalf("Failed to read destination file: %v", err)
	}
	if string(dstData) != string(testData) {
		t.Errorf("Data mismatch, Expected: %q, Got: %q", string(testData), string(dstData))
	}

	// Clean up the destination file
	if err := os.Remove(dst); err != nil {
		t.Fatalf("Failed to remove destination file: %v", err)
	}
}

func TestCopyDir(t *testing.T) {
	src := "src"
	dst := "dst"

	// Create source directory and files
	err := os.MkdirAll(src+"/subdir", 0775)
	if err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}
	err = os.WriteFile(src+"/file1.txt", []byte("Hello, World!"), 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}
	err = os.WriteFile(src+"/subdir/file2.txt", []byte("Another test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create source file: %v", err)
	}

	// Test copy
	err = CopyDir(dst, src)
	if err != nil {
		t.Fatalf("CopyDir failed: %v", err)
	}

	// Check if files and directories are copied
	entries, err := os.ReadDir(dst)
	if err != nil {
		t.Fatalf("Failed to read destination directory: %v", err)
	}
	if len(entries) != 2 {
		t.Errorf("Expected 2 entries in destination, got %d", len(entries))
	}

	// Check files content
	content1, err := os.ReadFile(dst + "/file1.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if string(content1) != "Hello, World!" {
		t.Errorf("Expected 'Hello, World!', got '%s'", string(content1))
	}

	content2, err := os.ReadFile(dst + "/subdir/file2.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}
	if string(content2) != "Another test" {
		t.Errorf("Expected 'Another test', got '%s'", string(content2))
	}

	// Clean up
	os.RemoveAll(src)
	os.RemoveAll(dst)
}

func TestCopyDirFailure(t *testing.T) {
	src := "src"
	dst := "/root"

	// Create source directory
	err := os.MkdirAll(src, 0775)
	if err != nil {
		t.Fatalf("Failed to create source directory: %v", err)
	}

	// Test copy to non-writable destination
	err = CopyDir(dst, src)
	if err == nil {
		t.Error("Expected error when copying to non-writable directory, got nil")
	}

	// Clean up
	os.RemoveAll(src)
}

func TestIsDir(t *testing.T) {
	// Test case 1: Directory exists
	dirPath := "/tmp"
	if !IsDir(dirPath) {
		t.Errorf("Expected %s to be a directory", dirPath)
	}

	// Test case 2: File exists but is not a directory
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if IsDir(tempFile.Name()) {
		t.Errorf("Expected %s to be a file, not a directory", tempFile.Name())
	}

	// Test case 3: Path does not exist
	nonExistentPath := "/nonexistent/path"
	if IsDir(nonExistentPath) {
		t.Errorf("Expected %s to not exist", nonExistentPath)
	}
}

func TestRemoveFile(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Ensure the file exists before attempting to remove it
	if _, err := os.Stat(tempFile.Name()); os.IsNotExist(err) {
		t.Fatalf("File does not exist: %s", tempFile.Name())
	}

	// Test removing the file
	if err := RemoveFile(tempFile.Name()); err != nil {
		t.Errorf("RemoveFile failed: %v", err)
	}

	// Ensure the file no longer exists after removal
	if _, err := os.Stat(tempFile.Name()); !os.IsNotExist(err) {
		t.Errorf("File still exists after removal: %s", tempFile.Name())
	}
}

func TestClearFile(t *testing.T) {
	// Create a temporary file with some content
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

	// Clear the file
	err = ClearFile(tempFile.Name())
	if err != nil {
		t.Fatalf("ClearFile failed: %v", err)
	}

	// Check if the file is empty
	fileInfo, err := os.Stat(tempFile.Name())
	if err != nil {
		t.Fatalf("Stat file failed: %v", err)
	}
	if fileInfo.Size() != 0 {
		t.Errorf("Expected file size to be 0, got: %d", fileInfo.Size())
	}

	// Try to read the file content
	content, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Read file failed: %v", err)
	}
	if len(content) != 0 {
		t.Errorf("Expected empty content, got: %s", content)
	}
}

func TestReadFileToString(t *testing.T) {
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

	content, err := ReadFileToString(tempFile.Name())
	if err != nil {
		t.Fatalf("ReadFileToString failed: %v", err)
	}

	if content != string(testData) {
		t.Errorf("Read content error, Expected: %q, Real: %q", string(testData), content)
	}

	// Test with non-existent file
	_, err = ReadFileToString("/nonexistent/path/to/file.txt")
	if err == nil {
		t.Error("Expected error when reading non-existent file, got nil")
	}
}

func TestReadFileByLine(t *testing.T) {
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

	lines, err := ReadFileByLine(tempFile.Name())
	if err != nil {
		t.Fatalf("ReadFileByLine failed: %v", err)
	}

	expectedLines := []string{"Hello, World!", "This is a test.", "Another line."}
	if !reflect.DeepEqual(lines, expectedLines) {
		t.Errorf("ReadFileByLine error, Expected: %v, Got: %v", expectedLines, lines)
	}
}

func TestReadFileByLine_EmptyFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatalf("Create temp file failed: %v", err)
	}
	defer os.Remove(tempFile.Name())

	lines, err := ReadFileByLine(tempFile.Name())
	if err != nil {
		t.Fatalf("ReadFileByLine failed: %v", err)
	}

	if len(lines) != 0 {
		t.Errorf("ReadFileByLine error, Expected empty file, Got: %v", lines)
	}
}

func TestReadFileByLine_FileNotFound(t *testing.T) {
	_, err := ReadFileByLine("/nonexistent/file.txt")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestFilesCurDir(t *testing.T) {
	// Test case 1: Directory does not exist
	vfn, err := FilesCurDir("/nonexistentdir")
	if vfn != nil || err != nil {
		t.Errorf("Expected no error when directory does not exist, got: %v", err)
	}

	// Test case 2: Directory exists but is empty
	tempDir, err := os.MkdirTemp("", "emptydir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	files, err := FilesCurDir(tempDir)
	if err != nil {
		t.Errorf("Expected no error when directory is empty, got: %v", err)
	}
	if len(files) != 0 {
		t.Errorf("Expected 0 files in empty directory, got: %d", len(files))
	}

	// Test case 3: Directory exists with files and subdirectories
	tempDir, err = os.MkdirTemp("", "testdir")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	testFile := filepath.Join(tempDir, "testfile.txt")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to write test file: %v", err)
	}

	subdir := filepath.Join(tempDir, "subdir")
	if err := os.Mkdir(subdir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}

	files, err = FilesCurDir(tempDir)
	if err != nil {
		t.Errorf("Expected no error when directory has files and subdirs, got: %v", err)
	}
	expectedFiles := []string{"testfile.txt"}
	if !reflect.DeepEqual(files, expectedFiles) {
		t.Errorf("Expected files %v, got: %v", expectedFiles, files)
	}
}

func TestIsZipFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     bool
	}{
		{
			name:     "Valid Zip File",
			filePath: "test_data/test.zip",
			want:     true,
		},
		{
			name:     "Invalid Zip File",
			filePath: "test_data/test_1.txt",
			want:     false,
		},
		{
			name:     "Non-Existent File",
			filePath: "/path/to/non/existent/file.zip",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsZipFile(tt.filePath); got != tt.want {
				t.Errorf("IsZipFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestZip_File(t *testing.T) {
	dst := "test_data/test_1.zip"
	src := "test_data/test_1.txt"

	err := Zip(dst, src)
	if err != nil {
		t.Errorf("Zip failed: %v", err)
	}
	defer RemoveFile(dst)

	if !IsZipFile(dst) {
		t.Errorf("Zip failed: %v", err)
	}
}

func TestZip_Directory(t *testing.T) {
	dst := "test_data.zip"
	src := "test_data"

	err := Zip(dst, src)
	if err != nil {
		t.Errorf("Zip failed: %v", err)
	}
	defer RemoveFile(dst)

	if !IsZipFile(dst) {
		t.Errorf("Zip failed: %v", err)
	}
}

func TestUnZip(t *testing.T) {
	// Setup
	src := "test_data/subdir"
	dst := "test_data/subdir.zip"

	// Create a test zip file
	err := Zip(dst, src)
	if err != nil {
		t.Fatalf("Zip failed: %v", err)
	}

	// Test unzipping the file
	err = UnZip(src+"_test", dst)
	if err != nil {
		t.Fatalf("UnZip failed: %v", err)
	}

	// Cleanup
	os.RemoveAll(src + "_test")
	RemoveFile(dst)
}

func TestZipAppendEntry(t *testing.T) {
	// Setup
	dst := "test_1.zip"
	src := "test_data/test.zip"

	defer os.Remove(dst)

	// Test appending directory to zip
	err := ZipAppendEntry(dst, src)
	if err != nil {
		t.Fatalf("ZipAppendEntry failed: %v", err)
	}

	// Verify the zip file content
	zipReader, err := zip.OpenReader(dst)
	if err != nil {
		t.Fatalf("Failed to open zip file: %v", err)
	}
	defer zipReader.Close()

	// Check the number of files in the zip
	if len(zipReader.File) != 3 { // Assuming testdir contains two files
		t.Errorf("Expected 2 files in zip, got %d", len(zipReader.File))
	}

	// Check the names of the files in the zip
	expectedNames := map[string]bool{
		"template-go":            false,
		"test.zip":               false,
		"__MACOSX/._template-go": false,
	}
	for _, file := range zipReader.File {
		if _, ok := expectedNames[file.Name]; ok {
			expectedNames[file.Name] = true
		} else {
			t.Errorf("Unexpected file in zip: %s", file.Name)
		}
	}

	// Check that all expected files were found
	for name, found := range expectedNames {
		if !found {
			t.Errorf("Expected file not found in zip: %s", name)
		}
	}
}

func TestIsLink(t *testing.T) {
	// Create a temporary directory for tests
	tempDir, err := os.MkdirTemp("", "testislink")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Test case 1: File is not a symlink
	tempFile, err := os.Create(filepath.Join(tempDir, "testfile"))
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	tempFile.Close()
	if IsLink(tempFile.Name()) {
		t.Errorf("Expected false for non-symlink file, got true")
	}

	// Test case 2: File is a symlink
	symlinkPath := filepath.Join(tempDir, "symlink")
	if err := os.Symlink(tempFile.Name(), symlinkPath); err != nil {
		t.Fatalf("Failed to create symlink: %v", err)
	}
	if !IsLink(symlinkPath) {
		t.Errorf("Expected true for symlink file, got false")
	}

	// Test case 3: Path does not exist
	nonExistentPath := filepath.Join(tempDir, "nonexistent")
	if IsLink(nonExistentPath) {
		t.Errorf("Expected false for non-existent path, got true")
	}
}

func TestFileMode(t *testing.T) {
	// Setup: Create a temporary directory and files
	tempDir, err := os.MkdirTemp("", "filemode_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) // Teardown

	// Create files with different permissions
	files := map[string]os.FileMode{
		"normal.txt": 0644,
		"dir":        0755,
		"link":       0644,
	}

	for name, mode := range files {
		path := filepath.Join(tempDir, name)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			t.Fatalf("Failed to create directory for %s: %v", name, err)
		}
		if err := os.WriteFile(path, []byte{}, mode); err != nil {
			t.Fatalf("Failed to create file %s: %v", name, err)
		}
		if name == "link" {
			linkPath := filepath.Join(tempDir, "link_to_normal.txt")
			if err := os.Symlink("normal.txt", linkPath); err != nil {
				t.Fatalf("Failed to create symlink %s: %v", linkPath, err)
			}
		}
	}

	// Test cases
	tests := []struct {
		name     string
		path     string
		expected os.FileMode
		hasError bool
	}{
		{"Existing file", filepath.Join(tempDir, "normal.txt"), 0644, false},
		{"Existing directory", filepath.Join(tempDir, "dir"), 0755, false},
		{"Existing symlink", filepath.Join(tempDir, "link_to_normal.txt"), os.ModeSymlink | 0755, false},
		{"Non-existing file", filepath.Join(tempDir, "non_existing.txt"), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mode, err := FileMode(tt.path)
			if (err != nil) != tt.hasError {
				t.Errorf("FileMode() error = %v, hasError %v", err, tt.hasError)
				return
			}
			if mode != tt.expected {
				t.Errorf("FileMode() got = %v, want %v", mode, tt.expected)
			}
		})
	}
}

func TestCurrPath(t *testing.T) {
	path := CurrPath()

	if path == "" {
		t.Error("CurrPath() returned an empty string")
	}

	// Check if the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("Path returned by CurrPath() does not exist: %s", path)
	}

	t.Log(path)
}

func TestMiMeType(t *testing.T) {
	// Test case 1: Valid file path
	filePath := "test_data/test_1.txt"
	expectedMimeType := "application/octet-stream"
	mimeType := MiMeType(filePath)
	if mimeType != expectedMimeType {
		t.Errorf("MiMeType() = %v, want %v", mimeType, expectedMimeType)
	}

	// Test case 2: Invalid file path
	invalidFilePath := "nonexistent.txt"
	mimeType = MiMeType(invalidFilePath)
	if mimeType != "" {
		t.Errorf("MiMeType() = %v, want ''", mimeType)
	}

	// Test case 3: Valid *os.File
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()
	mimeType = MiMeType(file)
	if mimeType != expectedMimeType {
		t.Errorf("MiMeType() = %v, want %v", mimeType, expectedMimeType)
	}

	// Test case 5: Empty file
	emptyFile, err := os.CreateTemp("", "empty.txt")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(emptyFile.Name())
	emptyFile.Close()

	mimeType = MiMeType(emptyFile.Name())
	if mimeType != "" {
		t.Errorf("MiMeType() = %v, want ''", mimeType)
	}
}

func TestFileSize(t *testing.T) {
	// Test case 1: Valid file path
	filePath := "test_data/test_1.txt"
	expectedSize := int64(109)
	size, err := FileSize(filePath)
	if err != nil {
		t.Errorf("fileSize() error = %v", err)
		return
	}
	if size != expectedSize {
		t.Errorf("fileSize() = %v, want %v", size, expectedSize)
	}

	// Test case 2: Invalid file path
	invalidFilePath := "nonexistent.txt"
	_, err = FileSize(invalidFilePath)
	if err == nil {
		t.Error("fileSize() expected error, got nil")
		return
	}
}

func TestDirSize(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dirsize_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test files and directories
	createTestFile := func(name string, size int64) error {
		filePath := filepath.Join(tempDir, name)
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(make([]byte, size)); err != nil {
			return err
		}
		return nil
	}

	if err := createTestFile("file1.txt", 10); err != nil {
		t.Fatal(err)
	}
	if err := createTestFile("file2.txt", 25); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(filepath.Join(tempDir, "subdir"), 0755); err != nil {
		t.Fatal(err)
	}
	if err := createTestFile("subdir/file3.txt", 15); err != nil {
		t.Fatal(err)
	}

	if err := createTestFile("empty", 0); err != nil {
		t.Fatal(err)
	}

	// Test cases
	tests := []struct {
		name     string
		dir      string
		expected int64
		hasError bool
	}{
		{"Valid directory", tempDir, 50, false},
		{"Empty directory", filepath.Join(tempDir, "empty"), 0, false},
		{"Non-existent directory", filepath.Join(tempDir, "nonexistent"), 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size, err := DirSize(tt.dir)
			if (err != nil) != tt.hasError {
				t.Errorf("DirSize() error = %v, hasError %v", err, tt.hasError)
				return
			}
			if size != tt.expected {
				t.Errorf("DirSize() got = %v, want %v", size, tt.expected)
			}
		})
	}
}

func TestDirSize_PermissionError(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := os.MkdirTemp("", "dirsize_perm_test")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create a subdirectory with no read permissions
	subdir := filepath.Join(tempDir, "inaccessible")
	if err := os.Mkdir(subdir, 0000); err != nil {
		t.Fatal(err)
	}

	// Try to get the directory size
	_, err = DirSize(tempDir)
	if err == nil {
		t.Error("DirSize() expected error for inaccessible directory, got nil")
	} else if !errors.Is(err, fs.ErrPermission) {
		t.Errorf("Expected permission error, got: %v", err)
	}
}

func TestMTime(t *testing.T) {
	// Test case 1: Valid file path
	filePath := "test_data/test_1.txt"
	mtime, err := MTime(filePath)
	if err != nil {
		t.Errorf("MTime() error = %v", err)
		return
	}
	if mtime <= 0 {
		t.Errorf("MTime() returned invalid timestamp: %v", mtime)
	}

	// Test case 2: Invalid file path
	invalidFilePath := "nonexistent.txt"
	_, err = MTime(invalidFilePath)
	if err == nil {
		t.Error("MTime() expected error, got nil")
		return
	}

	// Test case 3: Compare with file info
	tempFile, err := os.CreateTemp("", "mtime_test")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())
	tempFile.Close()

	fileInfo, err := os.Stat(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to get file info: %v", err)
	}

	mtime, err = MTime(tempFile.Name())
	if err != nil {
		t.Errorf("MTime() error = %v", err)
		return
	}

	if mtime != fileInfo.ModTime().Unix() {
		t.Errorf("MTime() returned different timestamp than file info")
	}

	// Test case 4: File modified time changes
	time.Sleep(2 * time.Second) // Wait for file system timestamp resolution
	os.Chtimes(tempFile.Name(), time.Now(), time.Now())

	mtime2, err := MTime(tempFile.Name())
	if err != nil {
		t.Errorf("MTime() error = %v", err)
		return
	}

	if mtime2 <= mtime {
		t.Errorf("MTime() did not reflect file modification time change")
	}
}

func TestSHA(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		shaType  int
		want     string
		wantErr  bool
	}{
		{
			name:     "SHA1_Success",
			filename: "test_data/test_1.txt",
			shaType:  1,
			want:     "68b6bcb45c3ce3f8c9d335f9da399c206857857c",
			wantErr:  false,
		},
		{
			name:     "SHA256_Success",
			filename: "test_data/test_1.txt",
			shaType:  256,
			want:     "c318de96d9bfcf31bff7a71c4679768b1270b8868011cb7636162d29e7cc0621",
			wantErr:  false,
		},
		{
			name:     "SHA512_Success",
			filename: "test_data/test_1.txt",
			shaType:  512,
			want:     "d3675f4aca9c33c4d3126d893bbfdf66e03b3d2c9379f25b3a5ec15b92ac76610f5a76b764adbdbd794cfd4b70b71099d3ffeae6bb615c1b8c301c8416997bcd",
			wantErr:  false,
		},
		{
			name:     "SHA1_Default",
			filename: "test_data/test_1.txt",
			shaType:  0, // Default should be SHA1
			want:     "",
			wantErr:  true,
		},
		{
			name:     "Invalid_SHAType",
			filename: "test_data/test_1.txt",
			shaType:  999,
			want:     "",
			wantErr:  true,
		},
		{
			name:     "File_Not_Found",
			filename: "nonexistent.txt",
			shaType:  1,
			want:     "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SHA(tt.filename, tt.shaType)
			if (err != nil) != tt.wantErr {
				t.Errorf("SHA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SHA() got = %v, want %v", got, tt.want)
			}
		})
	}
}
