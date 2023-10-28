package main
import(
	"fmt"
)
type board struct {
	size int
	length int
	characters []rune
	neighbors []neighborIdxList
}

type neighborIdxList struct {
	indices []int
}

type neighborTile struct {
	char rune
	index int
}

func buildBoard(chars string, size int) *board {
	
	if size*size == len(chars) {
		b := new(board)
		b.size = size
		b.length = len(chars)
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

			for ci := range neighborCols {
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

// func getNeighborTiles
func (b *board) getNeighborTiles(idx int) []neighborTile {
	var nl []neighborTile

	
	if b.characters[idx] != ' ' {
		for _, ni := range b.neighbors[idx].indices {
			if b.characters[ni] != ' ' {
				var np neighborTile
				np.char = b.characters[ni]
				np.index = ni
				nl = append(nl, np)
			}
		}
	}
	return nl
}

// func removeWord
func (b *board) removeWord(indices []int) *board {
	var retBoard board

	retBoard.length = b.length
	retBoard.characters = append(retBoard.characters, b.characters...)
	retBoard.neighbors = b.neighbors
	retBoard.size = b.size

	// blank the used letters
	for _, i := range indices {
		retBoard.characters[i] = ' '
	}

	// collapse the grid down onto the blanks
	for i := retBoard.length - 1; i >= 0; i-- {
		if retBoard.characters[i] == ' ' {
			for j := i - retBoard.size; j >= 0; j -= retBoard.size {
				if retBoard.characters[j] != ' ' {
					retBoard.characters[i] = retBoard.characters[j]
					retBoard.characters[j] = ' '
					break
				}
			}
		}
	}

	return &retBoard
}

// func
func (b *board) printBoard(name string) {
	fmt.Println(name)
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fmt.Printf("%c", b.characters[i*b.size + j])
		}
		fmt.Println()
	}
}
