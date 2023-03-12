package main

type StateAction func(state *State, idx int)

// AttachLeft は左側の単語を右側の単語の親にします
func AttachLeft(state *State, idx int) {
	parent := state.pending[idx]
	child := state.pending[idx+1]

	sta