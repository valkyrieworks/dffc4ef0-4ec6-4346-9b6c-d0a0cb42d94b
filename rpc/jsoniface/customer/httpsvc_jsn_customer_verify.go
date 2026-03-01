package customer

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func VerifyHttpsvcCustomerCreateHttpsvcCaller(t *testing.T) {
	processor := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("REDACTED"))
	})
	ts := httptest.NewServer(processor)
	defer ts.Close()

	timestampTransportsec := httptest.NewTLSServer(processor)
	defer timestampTransportsec.Close()
	//
	//
	timestampTransportsec.Config.ErrorLog = log.New(io.Discard, "REDACTED", 0)

	for _, verifyWebroute := range []string{ts.URL, timestampTransportsec.URL} {
		u, err := freshProcessedWebroute(verifyWebroute)
		require.NoError(t, err)
		callProc, err := createHttpsvcCaller(verifyWebroute)
		require.Nil(t, err)

		location, err := callProc(u.Scheme, u.ObtainMachineUsingRoute())
		require.NoError(t, err)
		require.NotNil(t, location)
	}
}

func Verify_processedurl(t *testing.T) {
	type verify struct {
		url                  string
		anticipatedWebroute          string
		anticipatedMachineUsingRoute string
		anticipatedCallLocator  string
	}

	verifies := map[string]verify{
		"REDACTED": {
			url:                  "REDACTED",
			anticipatedWebroute:          "REDACTED",
			anticipatedMachineUsingRoute: "REDACTED",
			anticipatedCallLocator:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedWebroute:          "REDACTED",
			anticipatedMachineUsingRoute: "REDACTED",
			anticipatedCallLocator:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedWebroute:          "REDACTED",
			anticipatedMachineUsingRoute: "REDACTED",
			anticipatedCallLocator:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedWebroute:          "REDACTED",
			anticipatedMachineUsingRoute: "REDACTED",
			anticipatedCallLocator:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedWebroute:          "REDACTED",
			anticipatedMachineUsingRoute: "REDACTED",
			anticipatedCallLocator:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedWebroute:          "REDACTED",
			anticipatedMachineUsingRoute: "REDACTED",
			anticipatedCallLocator:  "REDACTED",
		},
	}

	for alias, tt := range verifies {
		//
		t.Run(alias, func(t *testing.T) {
			processed, err := freshProcessedWebroute(tt.url)
			require.NoError(t, err)
			require.Equal(t, tt.anticipatedCallLocator, processed.ObtainCallLocator())
			require.Equal(t, tt.anticipatedWebroute, processed.ObtainShortenedWebroute())
			require.Equal(t, tt.anticipatedMachineUsingRoute, processed.ObtainMachineUsingRoute())
		})
	}
}
