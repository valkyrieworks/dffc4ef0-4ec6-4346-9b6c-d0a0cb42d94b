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
func compressPath(src, target string) error {
	compressRecord, err := os.Create(target)
	if err != nil {
		return err
	}
	defer compressRecord.Close()

	compressPersistor := zip.NewWriter(compressRecord)
	defer compressPersistor.Close()

	pathAlias := filepath.Base(target)
	foundationPath := strings.TrimSuffix(pathAlias, filepath.Ext(pathAlias))

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
		heading.Name = filepath.Join(foundationPath, strings.TrimPrefix(route, src))

		//
		//
		if details.IsDir() {
			heading.Name += "REDACTED"
		} else {
			heading.Method = zip.Deflate
		}

		headingPersistor, err := compressPersistor.CreateHeader(heading)
		if err != nil {
			return err
		}

		if details.IsDir() {
			return nil
		}

		record, err := os.Open(route)
		if err != nil {
			return err
		}
		defer record.Close()

		_, err = io.Copy(headingPersistor, record)
		return err
	})
}

//
//
func duplicateRecord(src, target string) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	originatingRecord, err := os.Open(src)
	if err != nil {
		return err
	}
	defer originatingRecord.Close()

	targetRecord, err := os.Create(target)
	if err != nil {
		return err
	}
	defer targetRecord.Close()

	if _, err = io.Copy(targetRecord, originatingRecord); err != nil {
		return err
	}

	originatingDetails, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(target, originatingDetails.Mode())
}

//
//
//
func persistStatusJSNTowardRecord(status any, dir, recordname string) error {
	statusJSN, err := json.MarshalIndent(status, "REDACTED", "REDACTED")
	if err != nil {
		return fmt.Errorf("REDACTED", err)
	}

	return os.WriteFile(path.Join(dir, recordname), statusJSN, 0o600)
}
