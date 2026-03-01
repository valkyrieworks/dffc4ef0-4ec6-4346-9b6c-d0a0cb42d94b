package net

import (
	"net"
	"strings"
)

//
//
func Relate(schemaLocation string) (net.Conn, error) {
	schema, location := SchemeAlsoLocation(schemaLocation)
	link, err := net.Dial(schema, location)
	return link, err
}

//
//
//
func SchemeAlsoLocation(overhearLocation string) (string, string) {
	scheme, location := "REDACTED", overhearLocation
	fragments := strings.SplitN(location, "REDACTED", 2)
	if len(fragments) == 2 {
		scheme, location = fragments[0], fragments[1]
	}
	return scheme, location
}

//
//
//
func ObtainLiberateChannel() (int, error) {
	location, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("REDACTED", location)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
