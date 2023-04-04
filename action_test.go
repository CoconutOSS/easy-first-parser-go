
package main

import (
	"testing"
)

func TestAttachLeft(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		makeWord("ms.", "NNP", 1, 2),
		makeWord("hang", "NNP", 2, 3),
		makeWord("plays", "VBZ", 3, 0),
		makeWord("elianti", "NNP", 4, 3),
		makeWord(".", ".", 5, 3),
	)
	s := NewState(words)
	AttachLeft(s, 3)
	p, ok := s.arcs[4]
	if !ok || p != 3 {
		t.Error("parent's index must be 3")
	}

	AttachLeft(s, 3)
	p, ok = s.arcs[5]
	if !ok || p != 3 {
		t.Error("parent's index must be 3")
	}

	if len(s.pending) != 4 {