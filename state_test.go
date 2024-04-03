package main

import (
	"testing"
)

func TestDeletePending(t *testing.T) {
	words := make([]*Word, 0)
	words = append(words,
		makeRootWord(),
		m