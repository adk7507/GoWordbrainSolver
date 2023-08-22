package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	s "strings"
)

var dict map[string]struct{}

func main() {

	dict = readDictionaryFile()

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	var splitRequestStr = s.Split(s.ToLower(r.URL.Path[1:]), ":")
	fmt.Fprintln(w, len(splitRequestStr))

	var gridSize = Isqrt(len(splitRequestStr[0]))
	fmt.Fprintf(w, "Grid is %dx%d\n", gridSize, gridSize)

	var wordLengths = s.Split(splitRequestStr[1], ",")
	fmt.Fprintf(w, "Word lengths are %s\n", wordLengths)

	fmt.Fprintf(w, "Dict length is %d\n", len(dict))

}

func Isqrt(n int) int {
	return sort.Search(n, func(x int) bool { return x*x+2*x+1 > n })
}

func wordInDictionary() bool {

	return false
}

func readDictionaryFile() map[string]struct{} {
	file, err := os.Open("john.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	dictLines := make(map[string]struct{})

	for scanner.Scan() {
		dictLines[scanner.Text()] = struct{}{}
	}
	return dictLines
}
