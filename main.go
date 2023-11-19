package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	s "strings"
)

func main() {
	// Board
	chars := "zcatzzzzzzzsixzzzzzzzteaz"
	// chars = "helpzlzpzzzzzzzz"
	size := 5
	wordLengths := []int{3}

	// The dictionary
    dict := trieInit()
	dict.readDictionaryFile("english_cleaned.txt")

	// Build the game board
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

	// // Begin word search
	// rootWord := resultTreeNode {
	// 	word: "",
	// 	gridIndices: []int{-1,-1,-1,-1},
	// 	collapsedBoard: b,
	// 	sourceBoard: b,
	// }

	//foundWords := findWord2(dict.root, b, wordLengths[0], []int{}, []rune{})
	rtn := resultTreeNode {
		word: "ROOT",
		gridIndices: []int{-1,-1,-1,-1},
		collapsedBoard: b,
		sourceBoard: b,
		nextWords: nil,
	}
	findPhrase2(0, wordLengths, &rtn, dict)
	printSubtree(&rtn, "", 2)

	fmt.Print("main: ")
	// fmt.Println(foundWords)
	//findPhrase(dict.root, wordLengths, &rootWord, 0)


	// Print results - move to HTML
	// fmt.Println("Collapsing")
	// flattened := make(map[string]int)
	// collapseSubtree(&rootWord, "", 3, &flattened)

	// fmt.Println(flattened)

	// http.HandleFunc("/", HelloServer)
	// http.ListenAndServe(":8080", nil)
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



