package fileutil

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func addFileToArchive(archive *zip.Writer, src string) error {
	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = strings.TrimPrefix(path, filepath.Dir(src)+"/")

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
			writer, err := archive.CreateHeader(header)
			if err != nil {
				return err
			}
			file, err := os.Open(path) //nolint:gosec
			if err != nil {
				return err
			}
			defer file.Close()
			if _, err := io.Copy(writer, file); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func zipFile(dst, src string) error {
	zf, err := os.Create(dst) //nolint:gosec
	if err != nil {
		return err
	}
	defer zf.Close()

	archive := zip.NewWriter(zf)
	defer func() { _ = archive.Close() }()

	return addFileToArchive(archive, src)
}

func addFileToArchive2(w *zip.Writer, basePath, baseInZip string) error {
	files, err := os.ReadDir(basePath)
	if err != nil {
		return err
	}
	if !strings.HasSuffix(basePath, "/") {
		basePath = basePath + "/"
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := os.ReadFile(basePath + file.Name()) //nolint:gosec
			if err != nil {
				return err
			}

			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			newBase := basePath + file.Name() + "/"
			_ = addFileToArchive2(w, newBase, baseInZip+file.Name()+"/")
		}
	}

	return nil
}

func zipDir(dst, src string) error {
	outFile, err := os.Create(dst) //nolint:gosec
	if err != nil {
		return err
	}
	defer outFile.Close()

	w := zip.NewWriter(outFile)

	err = addFileToArchive2(w, src, "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func safeFilepathJoin(path1, path2 string) (string, error) {
	relPath, err := filepath.Rel(".", path2)
	if err != nil || strings.HasPrefix(relPath, "..") {
		return "", fmt.Errorf("(zipslip) filepath is unsafe %q: %v", path2, err)
	}
	if path1 == "" {
		path1 = "."
	}
	return filepath.Join(path1, filepath.Join("/", relPath)), nil
}

// escapeCSVField change `\"` to `\"\"` when field contains delimiter.
func escapeCSVField(field string, delimiter rune) string {
	// change `"` -> `""`
	escapeField := strings.ReplaceAll(field, "\"", "\"\"")

	// If field contains [delimiter, `\"`, `\n`], add sournding `"..."`
	if strings.ContainsAny(escapeField, string(delimiter)+"\"\n") {
		escapeField = fmt.Sprintf("\"%s\"", escapeField)
	}

	return escapeField
}

func isCsvSupportedType(val any) bool {
	switch val.(type) {
	case bool,
		rune, string,
		int, int64,
		float32, float64,
		uint, byte,
		complex128, complex64,
		uintptr:
		return true
	default:
		return false
	}
}
