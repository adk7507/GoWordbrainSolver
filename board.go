package main
import(
	"fmt"
	"slices"
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

func (b *board) getUnvisitedTiles(visitedIdcs []int) []int {
	tiles := []int{}
	for i := 0; i < b.length; i++ {
		if !slices.Contains(visitedIdcs, i) {
			tiles = append(tiles, i)
		}
	}
	return tiles
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

// 
func (b *board) printBoard(name string) {
	fmt.Println(name)
	for i := 0; i < b.size; i++ {
		for j := 0; j < b.size; j++ {
			fmt.Printf("%c", b.characters[i*b.size + j])
		}
		fmt.Println()
	}
}

func findPhrase( wordIndex int, wordLengths []int, parentWord *resultTreeNode, dictionary *dictionaryTrie) {

	gameBoard := parentWord.collapsedBoard
	foundWords := findWord(dictionary.root, gameBoard, wordLengths[wordIndex], []int{}, []rune{})

	if len(foundWords) > 0 {
		parentWord.nextWords = []*resultTreeNode{}
		for _, fw := range foundWords {
			rtn := resultTreeNode {
				word: fw.word,
				gridIndices: fw.tiles,
				collapsedBoard: fw.collapsedBoard,
				sourceBoard: gameBoard,
				nextWords: nil,
			}

			if wordIndex == len(wordLengths) -1 {
				parentWord.nextWords = append(parentWord.nextWords, &rtn)
			} else {
				findPhrase(wordIndex + 1, wordLengths, &rtn, dictionary)
				if rtn.nextWords != nil && len(rtn.nextWords) > 0 {
					parentWord.nextWords = append(parentWord.nextWords, &rtn)
				}
			}
			
		}

	}

// 	// get collapsed board from parent
// 	// find all words of spec len
// 	// generate NODE for each found word
// 	// if this is the last word
// 		// attach nodes to the parent
// 	// if this is not the last word
// 		// find next words for each found NODE
// 		// if next words are present in these NODEs
// 			// attach NODEs to the parent

}

type foundWord struct {
	tiles []int
	word string
	collapsedBoard *board
}

func findWord(parentChar *dictionaryTreeNode, gameBoard *board, wordLength int, visitedTiles []int, visitedChars []rune) ([]foundWord) {

	if len(visitedChars) == 0 {
		allFoundWords := []foundWord{}
		for i, c := range gameBoard.characters {
			if c == ' ' {
				continue
			} 
			nvt := make([]int, len(visitedTiles)+1)
			nvc := make([] rune, len(visitedChars)+1)
			copy(nvt, append(visitedTiles, i))
			copy(nvc, append(visitedChars, c))
			foundWords := findWord(parentChar.findChildNode(c), gameBoard, wordLength, nvt, nvc)
			if len(foundWords) > 0 {
				allFoundWords = append(allFoundWords, foundWords...)
			}
		}		
		return allFoundWords

	} else if wordLength == len(visitedChars) && parentChar.wordEnds {

		fw := foundWord{
			tiles: visitedTiles,
			word: string(visitedChars),
			collapsedBoard: gameBoard.removeWord(visitedTiles),
		}		
		return []foundWord{fw}

	}  else if wordLength > len(visitedChars) {
		allFoundWords := []foundWord{}
		neighborTiles := gameBoard.getNeighborTiles(visitedTiles[len(visitedTiles)-1])
		for _, nt := range neighborTiles  {
			if !slices.Contains(visitedTiles, nt.index) {
				if nextDTN := parentChar.findChildNode(nt.char); nextDTN != nil {
					nvt := make([]int, len(visitedTiles)+1)
					nvc := make([] rune, len(visitedChars)+1)
					copy(nvt, append(visitedTiles, nt.index))
					copy(nvc, append(visitedChars, nt.char))
					foundWords := findWord(nextDTN, gameBoard, wordLength, nvt, nvc)
					allFoundWords = append(allFoundWords, foundWords...)	
				}
			}
		}
		
		return allFoundWords
	}

	return []foundWord{}
}

