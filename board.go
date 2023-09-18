package main
import(
	"fmt"
)
type board struct {
	size int
	len int
	characters []rune
	neighbors []neighborIdxList
}

type neighborIdxList struct {
	indices []int
}

type neighborPiece struct {
	char rune
	index int
}

func buildBoard(chars string, size int) *board {
	
	if size*size == len(chars) {
		b := new(board)
		b.size = size
		b.len = len(chars)
		fmt.Println(chars)
		b.characters = []rune(chars)
		fmt.Println(b.characters)
		// connect neighbors
		for i := range chars {
			r := i / size
			c := i % size
	
			neighborRows := []int{r, r-1, r-1, r-1, r, r+1, r+1, r+1}
			neighborCols := []int{c+1, c+1, c, c-1, c-1, c-1, c, c+1}

			var currentNeighbors neighborIdxList

			for ci, _ := range neighborCols {
				if neighborCols[ci] >= 0 && neighborCols[ci] < b.size && neighborRows[ci] >= 0 && neighborRows[ci] < b.size {
					currentNeighbors.indices = append(currentNeighbors.indices, neighborRows[ci] * b.size + neighborCols[ci])
				}
			}

			b.neighbors = append(b.neighbors, currentNeighbors)
		}
		return b
	}
	return nil
}

// func getNeighborCharacters
func (b *board) getNeighborCharacters(idx int) []neighborPiece {
	var nl []neighborPiece

	
	if b.characters[idx] != ' ' {
		for _, ni := range b.neighbors[idx].indices {
			if b.characters[ni] != ' ' {
				var np neighborPiece
				np.char = b.characters[ni]
				np.index = ni
				nl = append(nl, np)
			}
		}
	}
	return nl
}

// func removeWord

// func

