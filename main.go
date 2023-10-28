package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"slices"
	"sort"
	s "strings"
)


func main() {
	// Board
	chars := "bpnieommaretvpisroamdotiv"
	size := 5
	wordLengths := []int{4, 7, 7, 7}



	// The dictionary of truth (and some extra "words")
    dict := trieInit()
	dict.readDictionaryFile("english_cleaned.txt")




	b := buildBoard(chars, size)
	if( b == nil) {
		fmt.Fprintf(os.Stderr, "buildBoard returned nil\nInput 1: %s\nInput 2: %d\n", chars, size)
		os.Exit(-1)
	}

	for i := range b.characters {
		fmt.Printf("%d - %c - ", i, b.characters[i])
		fmt.Println(b.neighbors[i].indices)
	}

	b.printBoard("Game Board: ")

	rootWord := searchResultTree {
		letters: "ROOT",
		indices: []int{-1,-1,-1,-1},
		collapsedBoard: b,
		sourceBoard: b,
	}

	// for j := 0; j < b.length; j++ {
	// 	findWordLetters(dict.root, j, []int{}, 3, []rune{}, &rootWord)
	// }

	findPhraseWords(dict.root, wordLengths, &rootWord, 0)

	printSubtree(&rootWord, "", len(wordLengths))


	// http.HandleFunc("/", HelloServer)
	// http.ListenAndServe(":8080", nil)
}

func printSubtree(word *searchResultTree, prev string, depth int) {
	newStr := fmt.Sprintf("%s %s", prev, word.letters)
	if depth == 0 {
		fmt.Printf("%d %s\n", depth, newStr)
	} else {
		for ci := 0; ci < len(word.nextWords); ci++ {
			printSubtree(word.nextWords[ci], newStr, depth - 1)
		}
	}
}

// func pruneResultTree(rootWord *boardWord, depth int) int {
// 	if rootWord.nextWords == nil {
// 		return 0
// 	}
// 	for i, w := range rootWord.nextWords {
// 		nextDepth := pruneResultTree(w, depth - 1)
// 		if nextDepth < depth - 1
// 	}

// }

func findPhraseWords(dictRoot *dictionaryNode, wordLenghts []int, parentWord *searchResultTree, level int) {

	if len(wordLenghts) == 0 {
		return
	}

	wordLength := wordLenghts[0]
	nextWordLengths := wordLenghts[1:]
	for ti := 0; ti < parentWord.collapsedBoard.length; ti++ {
		if parentWord.collapsedBoard.characters[ti] != ' ' {
			findWordLetters(dictRoot, ti, []int{}, wordLength, []rune{}, parentWord)
		}
	}



	for wi := 0; wi < len(parentWord.nextWords); wi++ {
		findPhraseWords(dictRoot, nextWordLengths, parentWord.nextWords[wi], level + 1)
	}

}

func findWordLetters(parentDictNode *dictionaryNode, currentTileIdx int, visitedTiles [] int, wordLength int, visitedChars [] rune, parentWord *searchResultTree) {

	gameBoard := parentWord.collapsedBoard
	currentDictNode := parentDictNode.findChildNode(parentWord.collapsedBoard.characters[currentTileIdx])
	if currentDictNode != nil {
		visitedTiles = append(visitedTiles, currentTileIdx)
		visitedChars = append(visitedChars, gameBoard.characters[currentTileIdx])
		neighborTiles := gameBoard.getNeighborTiles(currentTileIdx)
		for _, nt := range neighborTiles  {
			if wordLength > 1 {
				if !slices.Contains(visitedTiles, nt.index) {
					findWordLetters(currentDictNode, nt.index, visitedTiles, wordLength-1, visitedChars, parentWord)
				} 
			} else {
				if currentDictNode.wordEnds {
					// found a legal word. add status to search results
					//  - create string of the word
					//  - collapse a board
					word := string(visitedChars)
					collapsedBoard := gameBoard.removeWord(visitedTiles)
					
					thisWord := searchResultTree {
						letters: word,
						indices: visitedTiles,
						collapsedBoard: collapsedBoard,
						sourceBoard: gameBoard,
					}
					if parentWord.nextWords == nil {
						parentWord.nextWords = []*searchResultTree{}
					}
					parentWord.nextWords = append(parentWord.nextWords, &thisWord)

					// fmt.Printf("%s - %s\n", word, fmt.Sprint(visitedTiles))
					return
				}
			}				
		}
	} 
}


func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	var splitRequestStr = s.Split(s.ToLower(r.URL.Path[1:]), ":")
	fmt.Fprintln(w, len(splitRequestStr))

	var characterGrid = splitRequestStr[0]
	var gridSize = Isqrt(len(characterGrid))
	fmt.Fprintf(w, "Grid is %dx%d\n", gridSize, gridSize)

	for i := 0; i < gridSize*gridSize; i++ {
		if i % gridSize == 0 {
			fmt.Fprintln(w)
		}
		fmt.Fprintf(w, "%c", characterGrid[i])

	}

	fmt.Fprintln(w)

	var wordLengths = s.Split(splitRequestStr[1], ",")
	fmt.Fprintf(w, "Word lengths are %s\n", wordLengths)


}

func Isqrt(n int) int {
	return sort.Search(n, func(x int) bool { return x*x+2*x+1 > n })
}

// Builds a trie and a map for timing comparison later
func (t *dictionaryTrie) readDictionaryFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		t.insert(line)
	}
}

