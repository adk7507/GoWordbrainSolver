package main

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	s "strings"
	"slices"
)

func main() {
	// The dictionary
    dict := trieInit()
	dict.readDictionaryFile("tiny_english.txt")

	// Build the game board
	chars := "ruuasenererigusiltoubtsdm"
	size := 5
	wordLengths := []int{5,5,3,7,5}

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

	rtn := resultTreeNode {
		word: "",
		gridIndices: []int{},
		collapsedBoard: b,
		sourceBoard: b,
		nextWords: nil,
	}
	findPhrase2(0, wordLengths, &rtn, dict)
	// printSubtree(&rtn, "",len(wordLengths))

	flattened := make(map[string]int,0)
	collapseSubtree(&rtn, "", len(wordLengths), &flattened)
	answers := make([]string, len(flattened))
	i := 0
	for k,_ := range flattened {
		answers[i] = k
		i++
	}
	fmt.Printf("Before sort: %d\n", len(answers))
	slices.Sort(answers)
	fmt.Printf("After sort: %d\n", len(answers))
	f, _ := os.Create("out.txt")


	for _,a := range answers {
		fmt.Println(a)
		f.WriteString(fmt.Sprintf("%s\n", a))
	}
	f.Close()
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



