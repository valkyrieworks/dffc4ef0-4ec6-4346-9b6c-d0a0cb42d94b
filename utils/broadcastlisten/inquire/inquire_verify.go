package inquire_test

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/iface/kinds"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire"
	"github.com/valkyrieworks/dffc4ef0-4ec6-4346-9b6c-d0a0cb42d94b/utils/broadcastlisten/inquire/grammar"
)

var _ broadcastlisten.Inquire = (*inquire.Inquire)(nil)

//
//
//
//
//
//
//
//
//
var apiIncidents = map[string][]string{
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

var apiKindIncidents = []kinds.Incident{
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
	{
		Kind: "REDACTED",
		Properties: []kinds.IncidentProperty{
			{
				Key:   "REDACTED",
				Datum: "REDACTED",
			},
		},
	},
}

func VerifyLargeNumerals(t *testing.T) {
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
		incidents  map[string][]string
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

			got, err := c.Aligns(tc.incidents)
			if err != nil {
				t.Errorf("REDACTED",
					tc.s, tc.incidents, err)
			}
			if got != tc.aligns {
				t.Errorf("REDACTED",
					tc.s, tc.incidents, got, tc.aligns)
			}
		})
	}
}

func VerifyAssembledAligns(t *testing.T) {
	var (
		transferTime = "REDACTED"
		transferMoment = "REDACTED"
	)

	//
	verifyScenarios := []struct {
		s       string
		incidents  map[string][]string
		aligns bool
	}{
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED" + time.Now().Format(grammar.TimeLayout)),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED" + transferTime),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED" + transferTime),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED" + time.Now().Format(grammar.MomentLayout)),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED" + transferMoment),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED", "REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED", "REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			true,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},
		{
			"REDACTED",
			freshVerifyIncidents("REDACTED"),
			false,
		},

		//
		{
			"REDACTED",
			apiIncidents, true,
		},
		{
			"REDACTED",
			apiIncidents, true,
		},
		{
			"REDACTED",
			apiIncidents, false,
		},
		{
			"REDACTED",
			apiIncidents, true,
		},
		{
			"REDACTED",
			apiIncidents, false,
		},
		{
			"REDACTED",
			apiIncidents, false,
		},
		{
			"REDACTED",
			apiIncidents, false,
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

			got, err := c.Aligns(tc.incidents)
			if err != nil {
				t.Errorf("REDACTED",
					tc.s, tc.incidents, err)
			}
			if got != tc.aligns {
				t.Errorf("REDACTED",
					tc.s, tc.incidents, got, tc.aligns)
			}
		})
	}
}

func arrangeIncidents(incidents []kinds.Incident) []kinds.Incident {
	sort.Slice(incidents, func(i, j int) bool {
		if incidents[i].Kind == incidents[j].Kind {
			return incidents[i].Properties[0].Key < incidents[j].Properties[0].Key
		}
		return incidents[i].Kind < incidents[j].Kind
	})
	return incidents
}

func VerifyAugmentIncidents(t *testing.T) {
	augmented := inquire.AugmentIncidents(apiIncidents)
	bz, err := json.Marshal(arrangeIncidents(augmented))
	require.NoError(t, err)
	bz2, err := json.Marshal(arrangeIncidents(apiKindIncidents))
	require.NoError(t, err)
	if string(bz) != string(bz2) {
		t.Errorf("REDACTED", string(bz), string(bz2))
	}
}

func VerifyEveryAlignsEvery(t *testing.T) {
	incidents := freshVerifyIncidents(
		"REDACTED",
		"REDACTED",
		"REDACTED",
		"REDACTED",
	)
	tokens := make([]string, 0)
	for k := range incidents {
		tokens = append(tokens, k)
	}
	for _, key := range tokens {
		delete(incidents, key)
		align, err := inquire.All.Aligns(incidents)
		if err != nil {
			t.Errorf("REDACTED", err)
		} else if !align {
			t.Errorf("REDACTED", incidents)
		}
	}
}

//
//
func appendFreshVerifyIncident(incidents map[string][]string, s string) {
	fragments := strings.Split(s, "REDACTED")
	key := fragments[0]
	for _, kv := range fragments[1:] {
		k, v := partitionTokval(kv)
		k = key + "REDACTED" + k
		incidents[k] = append(incidents[k], v)
	}
}

//
//
func freshVerifyIncidents(ss ...string) map[string][]string {
	incidents := make(map[string][]string)
	for _, s := range ss {
		appendFreshVerifyIncident(incidents, s)
	}
	return incidents
}

func partitionTokval(s string) (key, datum string) {
	kv := strings.SplitN(s, "REDACTED", 2)
	return kv[0], kv[1]
}
