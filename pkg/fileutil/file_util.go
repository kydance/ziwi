// =============================================================================
/*!
 *  @file       file_util.go
 *  @brief      Package fileutil implements some basic file operations.
 *  @author     kydenlu
 *  @date       2024.09
 *  @note
 */
// =============================================================================

package fileutil

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// FileReader is a wrapper of bufio.Reader,
// supporting offset seek, reading one line at a time.
type FileReader struct {
	*bufio.Reader

	fil *os.File // pointer to the file
	off int64    // offset of file cursor
}

// NewFileReader creates a new FileReader with the given path.
func NewFileReader(path string) (*FileReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return &FileReader{
		Reader: bufio.NewReader(f),
		fil:    f,
		off:    0,
	}, nil
}

// ReadLine reads a line and returns it excluding the trailing '\r' and '\n'.
func (fr *FileReader) ReadLine() (string, error) {
	dat, err := fr.Reader.ReadBytes('\n')
	fr.off += int64(len(dat))
	if err == nil || err == io.EOF {
		for len(dat) > 0 && (dat[len(dat)-1] == '\r' || dat[len(dat)-1] == '\n') {
			dat = dat[:len(dat)-1]
		}
		return string(dat), err
	}

	return "", err
}

// Offset returns the current offset of the file cursor.
func (fr *FileReader) Offset() int64 { return fr.off }

// SeekOffset sets the file cursor to the given offset.
func (fr *FileReader) SeekOffset(off int64) error {
	_, err := fr.fil.Seek(off, 0)
	if err != nil {
		return err
	}

	fr.Reader = bufio.NewReader(fr.fil)
	fr.off = off
	return nil
}

// Close closes the opened file.
func (fr *FileReader) Close() error { return fr.fil.Close() }

// ------------------------------------------------------------

// IsExist checks if a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

// CreateFile creates a file with the given path.
func CreateFile(path string) bool {
	f, err := os.Create(path)
	if err != nil {
		return false
	}

	defer f.Close()
	return true
}

// CreateDir creates a directory (perm: 0777) with the given path.
//
//	dir: absolute path like `/dev/null`
func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// CopyFile copies a file from src to dst. Support relative / absolute path.
func CopyFile(dst, src string) error {
	sf, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer df.Close()

	_, err = io.Copy(df, sf)
	return err
}

// RemoveFile removes a specifice file.
func RemoveFile(file string) error {
	return os.Remove(file)
}

// ClearFile clears a file, that is, it will write an empty string to the file.
func ClearFile(file string) error {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("")
	return err
}

// CopyDir copies a directory including all subdirectories and files from src to dst recursively.
// Support relative / absolute path. If dst does not exist, it will return error.
func CopyDir(dst, src string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to get source directory info: %w", err)
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("source is not a directory: %s", src)
	}

	err = os.MkdirAll(dst, 0775)
	if err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	vEntry, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	for _, entry := range vEntry {
		_src := filepath.Join(src, entry.Name())
		_dst := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err := CopyDir(_dst, _src)
			if err != nil {
				return err
			}
		} else {
			err := CopyFile(_dst, _src)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// IsDir checks if a path is a directory.
func IsDir(path string) bool {
	f, err := os.Stat(path)
	if err != nil {
		return false
	}
	return f.IsDir()
}
