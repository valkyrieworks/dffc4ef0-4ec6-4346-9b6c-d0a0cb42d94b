package texts

import (
	"fmt"
	"strings"
)

//
func TextInsideSection(a string, catalog []string) bool {
	for _, b := range catalog {
		if b == a {
			return true
		}
	}
	return false
}

//
//
//
//
//
func PartitionAlsoShave(s, sep, delimiters string) []string {
	if s == "REDACTED" {
		return []string{}
	}

	spl := strings.Split(s, sep)
	for i := 0; i < len(spl); i++ {
		spl[i] = strings.Trim(spl[i], delimiters)
	}
	return spl
}

//
//
//
//
//
func PartitionAlsoShortenBlank(s, sep, delimiters string) []string {
	if s == "REDACTED" {
		return []string{}
	}

	spl := strings.Split(s, sep)
	unBlankTexts := make([]string, 0, len(spl))
	for i := 0; i < len(spl); i++ {
		component := strings.Trim(spl[i], delimiters)
		if component != "REDACTED" {
			unBlankTexts = append(unBlankTexts, component)
		}
	}
	return unBlankTexts
}

//
func EqualsCODETxt(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, b := range []byte(s) {
		if b < 32 || b > 126 {
			return false
		}
	}
	return true
}

//
func CODEShave(s string) string {
	r := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		switch {
		case b == 32:
			continue //
		case 32 < b && b <= 126:
			r = append(r, b)
		default:
			panic(fmt.Sprintf("REDACTED", b))
		}
	}
	return string(r)
}

//
func TextSectionEquivalent(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
