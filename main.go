package main

import (
	"fmt"
	"net/http"
	"sort"
	s "strings"
	"os"
	"bufio"
	"log"
)

var wordDict map[string]bool
func main() {

	// Dictionary tree
    // myTrie := trieInit()

	// wordDict = myTrie.readDictionaryFile("english_cleaned.txt")

    // words_Search := []string{"aqua", "pinkertony", "card", "care","cat", "dog","can"}
    // for  _, wr := range words_Search {
    //     found := myTrie.search(wr)
    //     if found == 1 {
    //         fmt.Printf("\"%s\" found in trie\n", wr)
    //     } else {
    //         fmt.Printf(" \"%s\" NOT found in trie\n", wr)
    //     }
    // }

	// Board
	b := buildBoard([]rune{'a','b','c','d','e','f','g','h','i'}, 3)
	for i := range b.characters {
		fmt.Printf("%d - %c - ", i, b.characters[i])
		fmt.Println(b.neighbors[i].indices)
	}

	fmt.Printf("%c neighbors: \n", b.characters[0])
	npl := b.getNeighborCharacters(0)
	for _, np := range(npl) {
		fmt.Printf("%d - %c\n", np.index, np.char)
	}

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

	fmt.Fprintf(w, "Dict length is %d\n", len(wordDict))

}

func Isqrt(n int) int {
	return sort.Search(n, func(x int) bool { return x*x+2*x+1 > n })
}

// Builds a trie and a map for timing comparison later
func (t *trie) readDictionaryFile(filename string) map[string]bool {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dictLines := make(map[string]bool)

	var line string
	for scanner.Scan() {
		line = scanner.Text()
		dictLines[line] = true
		t.insert(line)
	}
	return dictLines
}

