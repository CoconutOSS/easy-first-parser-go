package main

import (
	"testing"
)

func TestDependencyAccuracy(t *testing.T) {
	g1 := []int{1, 2, 3}
	g2 := []int{1, 2, 3}
	g3 := []int{1, 2, 3, 4}
	g := [][]int{g1, g2, g3}

