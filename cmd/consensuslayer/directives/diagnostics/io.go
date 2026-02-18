package diagnostics

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//
//
//
func zipDir(src, dest string) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	dirName := filepath.Base(dest)
	baseDir := strings.TrimSuffix(dirName, filepath.Ext(dirName))

	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		//
		//
		header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, src))

		//
		//
		if info.IsDir() {
			header.Name += "REDACTED"
		} else {
			header.Method = zip.Deflate
		}

		headerWriter, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(headerWriter, file)
		return err
	})
}

//
//
func copyFile(src, dest string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	if _, err = io.Copy(destFile, srcFile); err != nil {
		return err
	}

	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dest, srcInfo.Mode())
}

//
//
//
func writeStateJSONToFile(state any, dir, filename string) error {
	stateJSON, err := json.MarshalIndent(state, "REDACTED", "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return os.WriteFile(path.Join(dir, filename), stateJSON, 0o600)
}
