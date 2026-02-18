package customer

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	kinds "github.com/valkyrieworks/rpc/jsonrpc/kinds"
)

const (
	//
	URICustomerQueryUID = kinds.JsonrpcIntegerUID(-1)
)

//
//
//
//
type URICustomer struct {
	location string
	customer  *http.Client
}

var _ HTTPCustomer = (*URICustomer)(nil)

//
//
//
func NewLocator(external string) (*URICustomer, error) {
	analyzedURL, err := newAnalyzedURL(external)
	if err != nil {
		return nil, err
	}

	httpCustomer, err := StandardHTTPCustomer(external)
	if err != nil {
		return nil, err
	}

	analyzedURL.CollectionStandardPlanHTTP()

	uriCustomer := &URICustomer{
		location: analyzedURL.FetchClippedURL(),
		customer:  httpCustomer,
	}

	return uriCustomer, nil
}

//
func (c *URICustomer) Invoke(ctx context.Context, procedure string,
	options map[string]any, outcome any,
) (any, error) {
	items, err := argsToURLItems(options)
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

	return unserializeReplyOctets(replyOctets, URICustomerQueryUID, outcome)
}
