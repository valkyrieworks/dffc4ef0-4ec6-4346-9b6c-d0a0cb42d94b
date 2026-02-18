package diagnose

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
func zipFolder(src, target string) error {
	zipEntry, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipEntry.Close()

	zipRecorder := zip.NewWriter(zipEntry)
	defer zipRecorder.Close()

	folderLabel := filepath.Base(target)
	rootFolder := strings.TrimSuffix(folderLabel, filepath.Ext(folderLabel))

	return filepath.Walk(src, func(route string, details os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		heading, err := zip.FileInfoHeader(details)
		if err != nil {
			return err
		}

		//
		//
		heading.Name = filepath.Join(rootFolder, strings.TrimPrefix(route, src))

		//
		//
		if details.IsDir() {
			heading.Name += "REDACTED"
		} else {
			heading.Method = zip.Deflate
		}

		headingRecorder, err := zipRecorder.CreateHeader(heading)
		if err != nil {
			return err
		}

		if details.IsDir() {
			return nil
		}

		entry, err := os.Open(route)
		if err != nil {
			return err
		}
		defer entry.Close()

		_, err = io.Copy(headingRecorder, entry)
		return err
	})
}

//
//
func cloneEntry(src, target string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	originEntry, err := os.Open(src)
	if err != nil {
		return err
	}
	defer originEntry.Close()

	targetEntry, err := os.Create(target)
	if err != nil {
		return err
	}
	defer targetEntry.Close()

	if _, err = io.Copy(targetEntry, originEntry); err != nil {
		return err
	}

	originDetails, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(target, originDetails.Mode())
}

//
//
//
func recordStatusJSONToEntry(status any, dir, filename string) error {
	statusJSON, err := json.MarshalIndent(status, "REDACTED", "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return os.WriteFile(path.Join(dir, filename), statusJSON, 0o600)
}
