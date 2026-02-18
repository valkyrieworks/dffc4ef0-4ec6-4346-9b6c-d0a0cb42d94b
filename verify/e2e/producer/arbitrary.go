package main

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
func groupings(items map[string][]any) []map[string]any {
	var keys []string //
	for key := range items {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return assembler(map[string]any{}, keys, items)
}

//
func assembler(front map[string]any, awaiting []string, items map[string][]any) []map[string]any {
	if len(awaiting) == 0 {
		return []map[string]any{front}
	}
	key, awaiting := awaiting[0], awaiting[1:]

	var outcome []map[string]any
	for _, item := range items[key] {
		route := map[string]any{}
		maps.Copy(route, front)
		route[key] = item
		outcome = append(outcome, assembler(route, awaiting, items)...)
	}
	return outcome
}

//
type consistentOption []any

func (uc consistentOption) Select(r *rand.Rand) any {
	return uc[r.Intn(len(uc))]
}

//
type likelihoodCollectionOption map[string]float64

func (pc likelihoodCollectionOption) Select(r *rand.Rand) []string {
	options := []string{}
	for item, likelihood := range pc {
		if r.Float64() <= likelihood {
			options = append(options, item)
		}
	}
	return options
}

//
type consistentCollectionOption []string

func (usc consistentCollectionOption) Select(r *rand.Rand) []string {
	var options []string //
	listings := r.Perm(len(usc))
	if len(listings) > 1 {
		listings = listings[:1+r.Intn(len(listings)-1)]
	}
	for _, i := range listings {
		options = append(options, usc[i])
	}
	return options
}

//
type scaledOption map[any]uint

func (wc scaledOption) Select(r *rand.Rand) any {
	sum := 0
	options := make([]any, 0, len(wc))
	for option, magnitude := range wc {
		sum += int(magnitude)
		options = append(options, option)
	}

	rem := r.Intn(sum)
	for _, option := range options {
		rem -= int(wc[option])
		if rem <= 0 {
			return option
		}
	}

	return nil
}
