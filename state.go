package main

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
)

type FvCache map[string][]int

type State struct