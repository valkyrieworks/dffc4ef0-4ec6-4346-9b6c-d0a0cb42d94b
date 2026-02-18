package os

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/valkyrieworks/utils/log"
)

type tracer interface {
	Details(msg string, keyvalues ...any)
}

//
//
func InterceptAlert(tracer tracer, cb func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			tracer.Details("REDACTED", "REDACTED", log.NewIdleFormat("REDACTED", sig))
			if cb != nil {
				cb()
			}
			os.Exit(0)
		}
	}()
}

//
func Abort() error {
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
func AssureFolder(dir string, style os.FileMode) error {
	err := os.MkdirAll(dir, style)
	if err != nil {
		return fmt.Errorf("REDACTED", dir, err)
	}
	return nil
}

func EntryPresent(entryRoute string) bool {
	_, err := os.Stat(entryRoute)
	return !os.IsNotExist(err)
}

func ReadEntry(entryRoute string) ([]byte, error) {
	return os.ReadFile(entryRoute)
}

func ShouldReaderEntry(entryRoute string) []byte {
	entryOctets, err := os.ReadFile(entryRoute)
	if err != nil {
		Quit(fmt.Sprintf("REDACTED", err))
		return nil
	}
	return entryOctets
}

func RecordEntry(entryRoute string, elements []byte, style os.FileMode) error {
	return os.WriteFile(entryRoute, elements, style)
}

func ShouldRecordEntry(entryRoute string, elements []byte, style os.FileMode) {
	err := RecordEntry(entryRoute, elements, style)
	if err != nil {
		Quit(fmt.Sprintf("REDACTED", err))
	}
}

//
func CloneEntry(src, dst string) error {
	sourceentry, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceentry.Close()

	details, err := sourceentry.Stat()
	if err != nil {
		return err
	}
	if details.IsDir() {
		return errors.New("REDACTED")
	}

	//
	destentry, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, details.Mode().Perm())
	if err != nil {
		return err
	}
	defer destentry.Close()

	_, err = io.Copy(destentry, sourceentry)
	return err
}
