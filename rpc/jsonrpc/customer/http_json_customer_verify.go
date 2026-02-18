package customer

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func VerifyHTTPCustomerCreateHTTPCaller(t *testing.T) {
	manager := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("REDACTED"))
	})
	ts := httptest.NewServer(manager)
	defer ts.Close()

	tsTLS := httptest.NewTLSServer(manager)
	defer tsTLS.Close()
	//
	//
	tsTLS.Config.ErrorLog = log.New(io.Discard, "REDACTED", 0)

	for _, verifyURL := range []string{ts.URL, tsTLS.URL} {
		u, err := newAnalyzedURL(verifyURL)
		require.NoError(t, err)
		callFn, err := createHTTPCaller(verifyURL)
		require.Nil(t, err)

		address, err := callFn(u.Scheme, u.FetchMachineWithRoute())
		require.NoError(t, err)
		require.NotNil(t, address)
	}
}

func Verify_analyzedurl(t *testing.T) {
	type verify struct {
		url                  string
		anticipatedURL          string
		anticipatedMachineWithRoute string
		anticipatedCallLocation  string
	}

	verifies := map[string]verify{
		"REDACTED": {
			url:                  "REDACTED",
			anticipatedURL:          "REDACTED",
			anticipatedMachineWithRoute: "REDACTED",
			anticipatedCallLocation:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedURL:          "REDACTED",
			anticipatedMachineWithRoute: "REDACTED",
			anticipatedCallLocation:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedURL:          "REDACTED",
			anticipatedMachineWithRoute: "REDACTED",
			anticipatedCallLocation:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedURL:          "REDACTED",
			anticipatedMachineWithRoute: "REDACTED",
			anticipatedCallLocation:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedURL:          "REDACTED",
			anticipatedMachineWithRoute: "REDACTED",
			anticipatedCallLocation:  "REDACTED",
		},

		"REDACTED": {
			url:                  "REDACTED",
			anticipatedURL:          "REDACTED",
			anticipatedMachineWithRoute: "REDACTED",
			anticipatedCallLocation:  "REDACTED",
		},
	}

	for label, tt := range verifies {
		//
		t.Run(label, func(t *testing.T) {
			analyzed, err := newAnalyzedURL(tt.url)
			require.NoError(t, err)
			require.Equal(t, tt.anticipatedCallLocation, analyzed.FetchCallLocation())
			require.Equal(t, tt.anticipatedURL, analyzed.FetchClippedURL())
			require.Equal(t, tt.anticipatedMachineWithRoute, analyzed.FetchMachineWithRoute())
		})
	}
}
