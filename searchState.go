package main

type phrase struct {
	firstWords []*legalWord
}

type legalWord struct {
	letters string
	indices []int
	collapsedBoard *board
	sourceBoard *board
	nextWords []*legalWord
}