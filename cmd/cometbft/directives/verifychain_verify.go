package directives

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyIncreaseIP(t *testing.T) {
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		increaseIP(ip)
		anticipated := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)

		increaseIP(ip)
		anticipated = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)

		//
		for range 255 {
			increaseIP(ip)
		}
		anticipated = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)
	}

	//
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		increaseIP(ip)
		anticipated := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)
	}

	//
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		increaseIP(ip)
		anticipated := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)

		//
		for range 0xFFFF {
			increaseIP(ip)
		}
		anticipated = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)
	}
}
