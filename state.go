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
	retur