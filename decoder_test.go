package main

import (
	"testing"
)

func TestDecode(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.",