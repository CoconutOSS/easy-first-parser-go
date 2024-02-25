package main

import (
	"reflect"
	"testing"
)

func TestExtractHeads(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.", "NNP", 1, 2),
	