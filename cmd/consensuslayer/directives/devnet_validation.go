package directives

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrementIP(t *testing.T) {
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		incrementIP(ip)
		expected := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(expected), "REDACTED", ip, expected)

		incrementIP(ip)
		expected = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(expected), "REDACTED", ip, expected)

		//
		for range 255 {
			incrementIP(ip)
		}
		expected = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(expected), "REDACTED", ip, expected)
	}

	//
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		incrementIP(ip)
		expected := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(expected), "REDACTED", ip, expected)
	}

	//
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		incrementIP(ip)
		expected := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(expected), "REDACTED", ip, expected)

		//
		for range 0xFFFF {
			incrementIP(ip)
		}
		expected = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(expected), "REDACTED", ip, expected)
	}
}
