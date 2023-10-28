package main


type searchResultTree struct {
	letters string
	indices []int
	collapsedBoard *board
	sourceBoard *board
	nextWords []*searchResultTree
	fullPhrase bool
}