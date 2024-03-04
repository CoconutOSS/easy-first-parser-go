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