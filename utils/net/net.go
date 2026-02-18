package net

import (
	"net"
	"strings"
)

//
//
func Link(schemaAddress string) (net.Conn, error) {
	schema, location := ProtocolAndLocation(schemaAddress)
	link, err := net.Dial(schema, location)
	return link, err
}

//
//
//
func ProtocolAndLocation(acceptAddress string) (string, string) {
	protocol, location := "REDACTED", acceptAddress
	segments := strings.SplitN(location, "REDACTED", 2)
	if len(segments) == 2 {
		protocol, location = segments[0], segments[1]
	}
	return protocol, location
}

//
//
//
func FetchReleasePort() (int, error) {
	address, err := net.ResolveTCPAddr("REDACTED", "REDACTED")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("REDACTED", address)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
