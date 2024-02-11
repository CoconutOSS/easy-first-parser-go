
package main

import (
	"reflect"
	"testing"
)

func TestEdgeFor(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ", 2, 1),
	)
	s := NewState(words)
	pair, err := EdgeFor(s, 0, 0)
	if err != nil {
		t.Error("error should be nil")
	}
	if !reflect.DeepEqual(pair, []int{0, 1}) {
		t.Error("pair shoud be [0, 1] but: ", pair)
	}
}

func TestIsValidFalse(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ", 2, 1),
	)
	s := NewState(words)
	goldArcs := make(map[int][]int)
	goldArcs[-1] = []int{0}
	goldArcs[0] = []int{1}
	goldArcs[1] = []int{2}
	if IsValid(s, 0, 0, goldArcs) != false {