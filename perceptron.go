package main

import (
	"errors"
	"math"
	"reflect"
)

// GoldArcs returns map of parent => children
func GoldArcs(sent *Sentence) map[int][]int {
	result := make(map[i