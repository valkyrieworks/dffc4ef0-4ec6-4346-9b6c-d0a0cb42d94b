package customer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	kinds "github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/rpc/jsoniface/kinds"
)

const (
	//
	URLCustomerSolicitUUID = kinds.JsonifaceIntegerUUID(-1)
)

//
//
//
//
type URLCustomer struct {
	location string
	customer  *http.Client
}

var _ HttpsvcCustomer = (*URLCustomer)(nil)

//
//
//
func FreshURL(distant string) (*URLCustomer, error) {
	processedWebroute, err := freshProcessedWebroute(distant)
	if err != nil {
		return nil, err
	}

	httpsvcCustomer, err := FallbackHttpsvcCustomer(distant)
	if err != nil {
		return nil, err
	}

	processedWebroute.AssignFallbackProtocolHttpsvc()

	urlCustomer := &URLCustomer{
		location: processedWebroute.FetchShortenedWebroute(),
		customer:  httpsvcCustomer,
	}

	return urlCustomer, nil
}

//
func (c *URLCustomer) Invocation(ctx context.Context, procedure string,
	parameters map[string]any, outcome any,
) (any, error) {
	items, err := argumentsTowardWebrouteItems(parameters)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.location+"REDACTED"+procedure,
		strings.NewReader(items.Encode()),
	)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	req.Header.Set("REDACTED", "REDACTED")

	reply, err := c.customer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}
	defer reply.Body.Close()

	replyOctets, err := io.ReadAll(reply.Body)
	if err != nil {
		return nil, fmt.Errorf("REDACTED", err)
	}

	return decodeReplyOctets(replyOctets, URLCustomerSolicitUUID, outcome)
}
