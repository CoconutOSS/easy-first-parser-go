
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