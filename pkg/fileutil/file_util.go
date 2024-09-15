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
	"archive/zip"
	"bufio"
	"bytes"

	// #nosec
	"crypto/sha1"

	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
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

// ReadFileToString reads a file and returns its content as a string.
func ReadFileToString(file string) (string, error) {
	vb, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	return string(vb), nil
}

// ReadFileByLine reads a file and returns its content line by line.
func ReadFileByLine(file string) ([]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	vs := make([]string, 0)
	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		sl := string(line)
		if err == io.EOF {
			break
		}

		if err != nil {
			continue
		}
		vs = append(vs, sl)
	}
	return vs, nil
}

// FilesCurDir returns all filenames in specific dir.
func FilesCurDir(dir string) ([]string, error) {
	if !IsExist(dir) {
		return nil, nil
	}

	vEntry, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	sz := len(vEntry)
	if sz == 0 {
		return nil, nil
	}

	vs := make([]string, 0)
	for _, entry := range vEntry {
		if IsDir(filepath.Join(dir, entry.Name())) {
			continue
		}
		vs = append(vs, entry.Name())
	}
	return vs, nil
}

// IsZipFile checks if a file is a zip file (PK\x03\x04).
func IsZipFile(file string) bool {
	f, err := os.Open(file)
	if err != nil {
		return false
	}
	defer f.Close()

	buf := make([]byte, 4)
	if n, err := f.Read(buf); err != nil || n < 4 {
		return false
	}

	return bytes.Equal(buf, []byte("PK\x03\x04"))
}

// Zip compresses a file or directory to a zip file.
// `dst` usually ends with ".zip".
func Zip(dst, src string) error {
	if IsDir(src) {
		return zipDir(dst, src)
	}
	return zipFile(dst, src)
}

// UnZip decompresses a zip file to dst.
// `src` usually ends with ".zip
func UnZip(dst, src string) error {
	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		path, err := safeFilepathJoin(dst, f.Name)
		if err != nil {
			return err
		}

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
		} else {
			err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
			if err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			// FIXME G110: Potential DoS vulnerability via decompression bomb (gosec)
			// #nosec
			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// ZipAppendEntry appends a file or directory to a zip file.
func ZipAppendEntry(dst, src string) error {
	tempFile, err := os.CreateTemp("", "temp.zip")
	if err != nil {
		return err
	}
	defer os.Remove(tempFile.Name())

	zipReader, err := zip.OpenReader(src)
	if err != nil {
		return err
	}

	archive := zip.NewWriter(tempFile)

	for _, item := range zipReader.File {
		itemReader, err := item.Open()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(item.FileInfo())
		if err != nil {
			return err
		}

		header.Name = item.Name
		targetItem, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		// FIXME G110: Potential DoS vulnerability via decompression bomb (gosec)
		// #nosec
		_, err = io.Copy(targetItem, itemReader)
		if err != nil {
			return err
		}
	}

	err = addFileToArchive(archive, src)
	if err != nil {
		return err
	}

	err = zipReader.Close()
	if err != nil {
		return err
	}

	err = archive.Close()
	if err != nil {
		return err
	}

	err = tempFile.Close()
	if err != nil {
		return err
	}

	return CopyFile(dst, tempFile.Name())
}

// IsLink checks if the specified path is symbolic link or not.
func IsLink(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode()&os.ModeSymlink != 0
}

// FileMode returns the file mode and permission of the specified file.
func FileMode(file string) (os.FileMode, error) {
	fi, err := os.Lstat(file)
	if err != nil {
		return 0, err
	}

	return fi.Mode(), nil
}

// CurrPath return current absolute path.
func CurrPath() string {
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		return filepath.Dir(filename)
	}

	return ""
}

// MiMeType returns file mime type of the specified file.
func MiMeType[T string | *os.File](file T) string {
	var mediatype string

	readBuffer := func(f *os.File) ([]byte, error) {
		buffer := make([]byte, 512)
		_, err := f.Read(buffer)
		if err != nil {
			return nil, err
		}
		return buffer, nil
	}

	switch fp := any(file).(type) {
	case string:
		f, err := os.Open(fp)
		if err != nil {
			return mediatype
		}
		buffer, err := readBuffer(f)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	case *os.File:
		buffer, err := readBuffer(fp)
		if err != nil {
			return mediatype
		}
		return http.DetectContentType(buffer)
	}

	return mediatype
}

// FileSize returns the given file size in bytes.
func FileSize(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	return f.Size(), nil
}

// DirSize walks the folder recusively and returns folder size in bytes.
func DirSize(dir string) (int64, error) {
	var size int64

	err := filepath.WalkDir(dir, func(_ string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				return err
			}

			size += info.Size()
		}
		return err
	})
	return size, err
}

// MTime return file modified time (Uinx timestamp).
func MTime(file string) (int64, error) {
	f, err := os.Stat(file)
	if err != nil {
		return 0, err
	}

	return f.ModTime().Unix(), nil
}

// SHA returns file SHA value.
//
//	SHATpye: [1, 256, 512]
func SHA(file string, SHAType ...int) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var h hash.Hash
	if len(SHAType) > 0 {
		switch SHAType[0] {
		case 1:
			// FIXME G401: Use of weak cryptographic primitive
			// #nosec
			h = sha1.New()

		case 256:
			h = sha256.New()

		case 512:
			h = sha512.New()
		default:
			return "", errors.New("param `SHAType` should be 1, 256 or 512")
		}
	} else {
		// FIXME G401: Use of weak cryptographic primitive
		// #nosec
		h = sha1.New()
	}

	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
