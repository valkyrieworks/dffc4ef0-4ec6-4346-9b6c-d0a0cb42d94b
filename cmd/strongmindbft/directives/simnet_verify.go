package directives

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func VerifyIncreaseINET(t *testing.T) {
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		increaseINET(ip)
		anticipated := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)

		increaseINET(ip)
		anticipated = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)

		//
		for range 255 {
			increaseINET(ip)
		}
		anticipated = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)
	}

	//
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		increaseINET(ip)
		anticipated := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)
	}

	//
	{
		ip := net.ParseIP("REDACTED")
		assert.NotNil(t, ip)

		increaseINET(ip)
		anticipated := net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)

		//
		for range 0xFFFF {
			increaseINET(ip)
		}
		anticipated = net.ParseIP("REDACTED")
		assert.True(t, ip.Equal(anticipated), "REDACTED", ip, anticipated)
	}
}
