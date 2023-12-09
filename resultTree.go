package main
import (
	"fmt"
)

type resultTreeNode struct {
	Word string
	gridIndices []int
	collapsedBoard *board
	sourceBoard *board
	NextWords []*resultTreeNode
}


func printSubtree(word *resultTreeNode, prev string, depth int) {
	newStr := fmt.Sprintf("%s %s", prev, word.Word)
	if depth == 0 {
		fmt.Printf("%d %s\n", depth, newStr)
	} else {
		for ci := 0; ci < len(word.NextWords); ci++ {
			printSubtree(word.NextWords[ci], newStr, depth - 1)
		}
	}
}

func collapseSubtree(word *resultTreeNode, prev string, depth int, flattened *map[string]int) {
	newStr := fmt.Sprintf("%s - %s", prev, word.Word)
	if depth == 0 {
		//fmt.Printf("%d %s\n", depth, newStr)
		(*flattened)[newStr] = depth
	} else {
		for ci := 0; ci < len(word.NextWords); ci++ {
			collapseSubtree(word.NextWords[ci], newStr, depth - 1, flattened)
		}
	}
}