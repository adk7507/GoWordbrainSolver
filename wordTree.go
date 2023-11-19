package main
import (
	"fmt"
)

type resultTreeNode struct {
	word string
	gridIndices []int
	collapsedBoard *board
	sourceBoard *board
	nextWords []*resultTreeNode
}


func printSubtree(word *resultTreeNode, prev string, depth int) {
	newStr := fmt.Sprintf("%s %s", prev, word.word)
	if depth == 0 {
		fmt.Printf("%d %s\n", depth, newStr)
	} else {
		for ci := 0; ci < len(word.nextWords); ci++ {
			printSubtree(word.nextWords[ci], newStr, depth - 1)
		}
	}
}

func collapseSubtree(word *resultTreeNode, prev string, depth int, flattened *map[string]int) {
	newStr :=  word.word
	if depth == 0 {
		//fmt.Printf("%d %s\n", depth, newStr)
		(*flattened)[newStr] = depth
	} else {
		for ci := 0; ci < len(word.nextWords); ci++ {
			collapseSubtree(word.nextWords[ci], newStr, depth - 1, flattened)
		}
	}
}