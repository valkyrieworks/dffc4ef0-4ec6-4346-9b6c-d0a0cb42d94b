package primary

import (
	"maps"
	"math/rand"
	"sort"
)

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
//
//
func permutations(elements map[string][]any) []map[string]any {
	var tokens []string //
	for key := range elements {
		tokens = append(tokens, key)
	}
	sort.Strings(tokens)
	return assembler(map[string]any{}, tokens, elements)
}

//
func assembler(header map[string]any, awaiting []string, elements map[string][]any) []map[string]any {
	if len(awaiting) == 0 {
		return []map[string]any{header}
	}
	key, awaiting := awaiting[0], awaiting[1:]

	var outcome []map[string]any
	for _, datum := range elements[key] {
		route := map[string]any{}
		maps.Copy(route, header)
		route[key] = datum
		outcome = append(outcome, assembler(route, awaiting, elements)...)
	}
	return outcome
}

//
type consistentOption []any

func (uc consistentOption) Select(r *rand.Rand) any {
	return uc[r.Intn(len(uc))]
}

//
type likelihoodAssignOption map[string]float64

func (pc likelihoodAssignOption) Select(r *rand.Rand) []string {
	selections := []string{}
	for record, likelihood := range pc {
		if r.Float64() <= likelihood {
			selections = append(selections, record)
		}
	}
	return selections
}

//
type consistentAssignOption []string

func (usc consistentAssignOption) Select(r *rand.Rand) []string {
	var selections []string //
	indices := r.Perm(len(usc))
	if len(indices) > 1 {
		indices = indices[:1+r.Intn(len(indices)-1)]
	}
	for _, i := range indices {
		selections = append(selections, usc[i])
	}
	return selections
}

//
type burdenedOption map[any]uint

func (wc burdenedOption) Select(r *rand.Rand) any {
	sum := 0
	selections := make([]any, 0, len(wc))
	for option, load := range wc {
		sum += int(load)
		selections = append(selections, option)
	}

	rem := r.Intn(sum)
	for _, option := range selections {
		rem -= int(wc[option])
		if rem <= 0 {
			return option
		}
	}

	return nil
}
