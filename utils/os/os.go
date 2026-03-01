package os

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/log"
)

type tracer interface {
	Details(msg string, tokvals ...any)
}

//
//
func EnsnareGesture(tracer tracer, cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			tracer.Details("REDACTED", "REDACTED", log.FreshIdleFormat("REDACTED", sig))
			if cb != nil {
				cb()
			}
			os.Exit(0)
		}
	}()
}

//
func Terminate() error {
	p, err := os.FindProcess(os.Getpid())
	if err != nil {
		return err
	}
	return p.Signal(syscall.SIGTERM)
}

func Quit(s string) {
	fmt.Println(s)
	os.Exit(1)
}

//
//
func AssurePath(dir string, style os.FileMode) error {
	err := os.MkdirAll(dir, style)
	if err != nil {
		return fmt.Errorf("REDACTED", dir, err)
	}
	return nil
}

func RecordPresent(recordRoute string) bool {
	_, err := os.Stat(recordRoute)
	return !os.IsNotExist(err)
}

func FetchRecord(recordRoute string) ([]byte, error) {
	return os.ReadFile(recordRoute)
}

func ShouldFetchRecord(recordRoute string) []byte {
	recordOctets, err := os.ReadFile(recordRoute)
	if err != nil {
		Quit(fmt.Sprintf("REDACTED", err))
		return nil
	}
	return recordOctets
}

func RecordRecord(recordRoute string, material []byte, style os.FileMode) error {
	return os.WriteFile(recordRoute, material, style)
}

func ShouldRecordRecord(recordRoute string, material []byte, style os.FileMode) {
	err := RecordRecord(recordRoute, material, style)
	if err != nil {
		Quit(fmt.Sprintf("REDACTED", err))
	}
}

//
func DuplicateRecord(src, dst string) error {
	sourcerecord, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourcerecord.Close()

	details, err := sourcerecord.Stat()
	if err != nil {
		return err
	}
	if details.IsDir() {
		return errors.New("REDACTED")
	}

	//
	destrecord, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, details.Mode().Perm())
	if err != nil {
		return err
	}
	defer destrecord.Close()

	_, err = io.Copy(destrecord, sourcerecord)
	return err
}
