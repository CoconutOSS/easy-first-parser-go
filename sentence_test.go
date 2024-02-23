package main

import (
	"reflect"
	"testing"
)

func TestExtractHeads(t *testing.T) {
	words := make([]*Word