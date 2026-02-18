package inquire_test

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/iface/kinds"
	"github.com/valkyrieworks/utils/broadcast"
	"github.com/valkyrieworks/utils/broadcast/inquire"
	"github.com/valkyrieworks/utils/broadcast/inquire/grammar"
)

var _ broadcast.Inquire = (*inquire.Inquire)(nil)

//
//
//
//
//
//
//
//
//
var apiEvents = map[string][]string{
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
	},
	"REDACTED": {
		"REDACTED",
	},
}

var apiKindEvents = []kinds.Event{
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.EventProperty{
			{
				Key:   "REDACTED",
				Item: "REDACTED",
			},
		},
	},
}

func VerifyLargeFigures(t *testing.T) {
	apiLargeCountVerify := map[string][]string{
		"REDACTED": {
			"REDACTED",
		},
		"REDACTED": {
			"REDACTED", //
		},
		"REDACTED": {
			"REDACTED",
		},
		"REDACTED": {
			"REDACTED", //
		},
	}

	verifyScenarios := []struct {
		s       string
		events  map[string][]string
		aligns bool
	}{
		//
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify, true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
		{
			"REDACTED",
			apiLargeCountVerify,
			true,
		},
	}

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i+1), func(t *testing.T) {
			c, err := inquire.New(tc.s)
			if err != nil {
				t.Fatalf("REDACTED", tc.s, err)
			}

			got, err := c.Aligns(tc.events)
			if err != nil {
				t.Errorf("REDACTED",
					tc.s, tc.events, err)
			}
			if got != tc.aligns {
				t.Errorf("REDACTED",
					tc.s, tc.events, got, tc.aligns)
			}
		})
	}
}

func VerifyBuiltAligns(t *testing.T) {
	var (
		transferDate = "REDACTED"
		transferTime = "REDACTED"
	)

	//
	verifyScenarios := []struct {
		s       string
		events  map[string][]string
		aligns bool
	}{
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED" + time.Now().Format(grammar.DateLayout)),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED" + transferDate),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED" + transferDate),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED" + time.Now().Format(grammar.TimeLayout)),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED" + transferTime),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED", "REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			newVerifyEvents("REDACTED"),
			false,
		},

		//
		{
			"REDACTED",
			apiEvents, true,
		},
		{
			"REDACTED",
			apiEvents, true,
		},
		{
			"REDACTED",
			apiEvents, false,
		},
		{
			"REDACTED",
			apiEvents, true,
		},
		{
			"REDACTED",
			apiEvents, false,
		},
		{
			"REDACTED",
			apiEvents, false,
		},
		{
			"REDACTED",
			apiEvents, false,
		},
	}

	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//

	for i, tc := range verifyScenarios {
		t.Run(fmt.Sprintf("REDACTED", i+1), func(t *testing.T) {
			c, err := inquire.New(tc.s)
			if err != nil {
				t.Fatalf("REDACTED", tc.s, err)
			}

			got, err := c.Aligns(tc.events)
			if err != nil {
				t.Errorf("REDACTED",
					tc.s, tc.events, err)
			}
			if got != tc.aligns {
				t.Errorf("REDACTED",
					tc.s, tc.events, got, tc.aligns)
			}
		})
	}
}

func arrangeEvents(events []kinds.Event) []kinds.Event {
	sort.Slice(events, func(i, j int) bool {
		if events[i].Kind == events[j].Kind {
			return events[i].Properties[0].Key < events[j].Properties[0].Key
		}
		return events[i].Kind < events[j].Kind
	})
	return events
}

func VerifyExtendEvents(t *testing.T) {
	extended := inquire.ExtendEvents(apiEvents)
	bz, err := json.Marshal(arrangeEvents(extended))
	require.NoError(t, err)
	bz2, err := json.Marshal(arrangeEvents(apiKindEvents))
	require.NoError(t, err)
	if string(bz) != string(bz2) {
		t.Errorf("REDACTED", string(bz), string(bz2))
	}
}

func VerifyAllAlignsAll(t *testing.T) {
	events := newVerifyEvents(
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	)
	keys := make([]string, 0)
	for k := range events {
		keys = append(keys, k)
	}
	for _, key := range keys {
		delete(events, key)
		align, err := inquire.All.Aligns(events)
		if err != nil {
			t.Errorf("REDACTED", err)
		} else if !align {
			t.Errorf("REDACTED", events)
		}
	}
}

//
//
func appendNewVerifyEvent(events map[string][]string, s string) {
	segments := strings.Split(s, "REDACTED")
	key := segments[0]
	for _, kv := range segments[1:] {
		k, v := divideObject(kv)
		k = key + "REDACTED" + k
		events[k] = append(events[k], v)
	}
}

//
//
func newVerifyEvents(ss ...string) map[string][]string {
	events := make(map[string][]string)
	for _, s := range ss {
		appendNewVerifyEvent(events, s)
	}
	return events
}

func divideObject(s string) (key, item string) {
	kv := strings.SplitN(s, "REDACTED", 2)
	return kv[0], kv[1]
}
