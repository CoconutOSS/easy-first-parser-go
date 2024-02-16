package main

type Sentence struct {
	words []*Word
}

// extract heads without root for evaluation
func (sent *