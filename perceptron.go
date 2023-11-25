package main

import (
	"errors"
	"math"
	"reflect"
)

// GoldArcs returns map of parent => children
func GoldArcs(sent *Sentence) map[int][]int {
	result := make(map[int][]int)
	for idx, w := range sent.words {
		head := w.head
		if children, ok := result[head]; ok {
			result[head] = append(children, idx)
		} else {
			result[head] = []int{idx}
		}
	}
	return result
}

// EdgeFor returns a pair of parent index and child index
func EdgeFor(state *State, actionID int, idx int) ([]int, error) {
	switch actionID {
	case 0:
		return []int{state.pending[idx].idx, state.pending[idx+1].idx}, nil
	case 1:
		return []int{state.pending[idx+1].idx, state.pending[idx].idx}, nil
	default:
		return nil, errors.New("Invalid line")
	}
}

// IsValid returns the chosen action/location pair is valid
func IsValid(state *State, actionID int, idx int, goldArcs map[int][]int) bool {
	pair, err := EdgeFor(state, actionID, idx)
	if err != nil {
		return false
	}
	pIdx := pair[0]
	cIdx := pair[1]
	containedInGoldArcs := false
	for _, i := range goldArcs[pIdx] {
		if cIdx == i {
			containedInGoldArcs = true
			break
		}
	}
	flag := false
	for _, cPrime := range goldArcs[cIdx] {
		if cIdx != state.arcs[cPrime] {
			flag = true
			break
		}
	}
	if !containedInGoldArcs || flag {
		return false
	}
	return true
}

type ActionIndexPair struct {
	action StateAction
	index  int
}

func (pair1 ActionIndexPair) SameActionIndexPair(pair2 ActionIndexPair) bool {
	return pair1.index == pair2.index &&
		reflect.ValueOf(pair1.action).Pointer() == reflect.ValueOf(pair2.action).Pointer()
}

func AllowedActions(state *State, goldArcs map[int][]int) []ActionIndexPair {
	result := make([]ActionIndexPair, 0)
	for actionID, f := range StateActions {
		for idx := 0; idx < len(state.pending)-1; idx++ {
			i