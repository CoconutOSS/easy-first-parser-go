package main

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
)

type FvCache map[string][]int

type State struct {
	pending []*Word
	arcs    map[int]int
	fvCache FvCache
}

func (state *State) cacheKeyStr(pair ActionIndexPair) string {
	funcName := runtime.FuncForPC(reflect.ValueOf(pair.action).Pointer()).Name()
	left := state.pending[pair.index]
	right := state.pending[pair.index+1]
	return funcName + ":" + strconv.Itoa(left.idx) + "-" + strconv.Itoa(right.idx)
}

func (state *State) InitFvCache() {
	for _, f := range StateActions {
		for idx := 0; idx < len(state.pending)-1; idx++ {
			pair := ActionIndexPair{f, idx}
			fv := ExtractFeatures(state, pair)
			state.fvCache[state.cacheKeyStr(pair)] = fv
		}
	}
}

func NewState(pending []*Word) *State {
	for _, w := range pending {
		w.children = make([]Word, 0)
	}
	p := ma