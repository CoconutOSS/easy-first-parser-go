package main

import (
	"testing"
)

func TestDeletePending(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.", "NNP", 0, -1),
		makeWord("hang", "NNP", 1, 0),
		makeWord("plays", "VBZ