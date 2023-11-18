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


func findPhrase(dictRoot *dictionaryTreeNode, wordLenghts []int, parentWord *resultTreeNode, level int) {

	if len(wordLenghts) == 0 {
		return
	}

	wordLength := wordLenghts[0]
	nextWordLengths := wordLenghts[1:]
	for ti := 0; ti < parentWord.collapsedBoard.length; ti++ {
		if parentWord.collapsedBoard.characters[ti] != ' ' {
			findWord(dictRoot, ti, []int{}, wordLength, []rune{}, parentWord)
		}
	}

	for wi := 0; wi < len(parentWord.nextWords); wi++ {
		findPhrase(dictRoot, nextWordLengths, parentWord.nextWords[wi], level + 1)
	}

}

func findWord(parentDictNode *dictionaryTreeNode, currentTileIdx int, visitedTiles [] int, wordLength int, visitedChars [] rune, parentWord *resultTreeNode) {

	gameBoard := parentWord.collapsedBoard
	currentDictNode := parentDictNode.findChildNode(parentWord.collapsedBoard.characters[currentTileIdx])
	if currentDictNode != nil {
		visitedTiles = append(visitedTiles, currentTileIdx)
		visitedChars = append(visitedChars, gameBoard.characters[currentTileIdx])
		neighborTiles := gameBoard.getNeighborTiles(currentTileIdx)
		for _, nt := range neighborTiles  {
			if wordLength > 1 {
				if !slices.Contains(visitedTiles, nt.index) {
					findWord(currentDictNode, nt.index, visitedTiles, wordLength-1, visitedChars, parentWord)
				} 
			} else {
				if currentDictNode.wordEnds {
					// found a legal word. add status to search results
					//  - create string of the word
					//  - collapse a board
					word := string(visitedChars)
					collapsedBoard := gameBoard.removeWord(visitedTiles)
					
					thisWord := resultTreeNode {
						word: word,
						gridIndices: visitedTiles,
						collapsedBoard: collapsedBoard,
						sourceBoard: gameBoard,
					}
					if parentWord.nextWords == nil {
						parentWord.nextWords = []*resultTreeNode{}
					}
					parentWord.nextWords = append(parentWord.nextWords, &thisWord)

					// fmt.Printf("%s - %s\n", word, fmt.Sprint(visitedTiles))
					return
				}
			}				
		}
	} 
}


// func findPhrase2( wordIndex int, wordLengths []int, parentWord *resultTreeNode, dictionary *dictionaryTreeNode) {
// 	wordLength := wordLengths[wordIndex]
// 	foundWords := findWord2(parentWord.collapsedBoard, wordLength)

// 	// get collapsed board from parent
// 	// find all words of spec len
// 	// generate NODE for each found word
// 	// if this is the last word
// 		// attach nodes to the parent
// 	// if this is not the last word
// 		// find next words for each found NODE
// 		// if next words are present in these NODEs
// 			// attach NODEs to the parent


// }

type foundWord struct {
	tiles []int
	word string
	collapsedBoard *board
}

func findWord2(parentChar *dictionaryTreeNode, gameBoard *board, wordLength int, visitedTiles []int, visitedChars []rune) ([]foundWord) {
	// if the word is finished, begin the return process
		// add the string word and visited tiles to the struct
	// if the word is not finished
		// get the neighbors
		// loop over the neighbors
		// find the next letters
		// append the returned words

	if len(visitedChars) == 0 {
		allFoundWords := []foundWord{}
		for i, c := range gameBoard.characters {
			foundWords := findWord2(parentChar.findChildNode(c), gameBoard, wordLength, append(visitedTiles, i), append(visitedChars, c))
			if len(foundWords) > 0 {
				allFoundWords = append(allFoundWords, foundWords...)
			}
		}
		// fmt.Println("visitedChars == 0")
		// fmt.Println(allFoundWords)
		return allFoundWords
	} else if wordLength == len(visitedChars) && parentChar.wordEnds {

		fw := foundWord{
			tiles: visitedTiles,
			word: string(visitedChars),
			collapsedBoard: gameBoard.removeWord(visitedTiles),
		}
		// fmt.Printf("wordLength %d == len(visitedChars) %d, visitedTiles: %s! ", wordLength, len(visitedChars), fmt.Sprint(visitedTiles))
		// fmt.Println(fw)
		return []foundWord{fw}
	}  else if wordLength > len(visitedChars) {
		allFoundWords := []foundWord{}
		neighborTiles := gameBoard.getNeighborTiles(visitedTiles[len(visitedTiles)-1])
		for _, nt := range neighborTiles  {
			if !slices.Contains(visitedTiles, nt.index) {
				if nextDTN := parentChar.findChildNode(nt.char); nextDTN != nil {
					newVisitedTiles :=  append(visitedTiles, nt.index)
					copyOfNewVisitedTiles := make([]int, len(newVisitedTiles))
					copy(copyOfNewVisitedTiles, newVisitedTiles)
					foundWords := findWord2(nextDTN, gameBoard, wordLength, copyOfNewVisitedTiles, append(visitedChars, nt.char))
					// fmt.Print("Special: ")
					// fmt.Printf("wordLength %d > len(visitedChars) %d, visitedTiles: %s! ", wordLength, len(visitedChars), fmt.Sprint(visitedTiles))
					
					// fmt.Println(foundWords)
					for _, fw := range foundWords {
						// fmt.Printf("\t%s\n", fmt.Sprint(fw))
						allFoundWords = append(allFoundWords, fw)
					}
					// if len(foundWords) > 0 {
					// 	allFoundWords = append(allFoundWords, foundWords...)
					// }
				}
			}
		}
		// fmt.Printf("wordLength %d > len(visitedChars) %d, visitedTiles: %s! ", wordLength, len(visitedChars), fmt.Sprint(visitedTiles))
		// fmt.Println(allFoundWords)
		return allFoundWords
	}
	return []foundWord{}

}


