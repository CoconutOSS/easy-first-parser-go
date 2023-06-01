
package main

import (
	"math"
	"reflect"
	"runtime"
	"strconv"
)

func NilSafePosTag(w *Word) string {
	posTag := ""
	if w != nil {
		posTag = w.posTag
	}
	return posTag
}

func NilSafePendingWord(state *State, idx int) *Word {
	if idx < 0 || idx >= len(state.pending) {
		return nil
	} else {
		return state.pending[idx]
	}
}

func addUnigramFeatures(features *[]int, state *State, actName string, idx int, prefix string) {
	if idx < 0 || idx >= len(state.pending) {
		return
	}
	w := state.pending[idx]
	lcp := NilSafePosTag(w.LeftMostChild())
	rcp := NilSafePosTag(w.RightMostChild())
	*features = append(*features,
		JenkinsHash(actName+"+"+prefix+"+surface:"+w.surface),
		JenkinsHash(actName+"+"+prefix+"+lemma:"+w.lemma),
		JenkinsHash(actName+"+"+prefix+"+posTag:"+w.posTag),
		JenkinsHash(actName+"+"+prefix+"+cposTag:"+w.cposTag),
		JenkinsHash(actName+"+"+prefix+"+posTag:"+w.posTag+"+leftmost:"+lcp),
		JenkinsHash(actName+"+"+prefix+"+posTag:"+w.posTag+"+rightmost:"+rcp),
		JenkinsHash(actName+"+"+prefix+"+posTag:"+w.posTag+"+leftmost:"+lcp+"+rightmost:"+rcp),
	)
}

func distStr(dist int) string {
	d := "0"
	switch dist {
	case 1:
		d = "1"
	case 2:
		d = "2"
	case 3:
		d = "3"
	case 4:
		d = "4"
	default:
		d = "5"
	}
	return d
}

func AddBigramFeatures(features *[]int, actName string, parent *Word, child *Word, prefix string) {
	if parent == nil || child == nil {
		return
	}

	plcp := NilSafePosTag(parent.LeftMostChild())
	prcp := NilSafePosTag(parent.RightMostChild())
	clcp := NilSafePosTag(child.LeftMostChild())
	crcp := NilSafePosTag(child.RightMostChild())

	*features = append(*features,
		JenkinsHash(actName+"+"+prefix+"+parent-surface:"+parent.surface+"+child-surface:"+child.surface),
		JenkinsHash(actName+"+"+prefix+"+parent-surface:"+parent.surface+"+child-posTag:"+child.posTag),
		JenkinsHash(actName+"+"+prefix+"+parent-posTag:"+parent.posTag+"+child-surface:"+child.surface),
		JenkinsHash(actName+"+"+prefix+"+parent-lemma:"+parent.lemma+"+child-lemma:"+child.lemma),