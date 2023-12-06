package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
	"context"
	"github.com/google/uuid"
	"strings"
	"regexp"
	"strconv"
)

var page *template.Template
var dict *dictionaryTrie

func init() {
	page, _ = template.ParseGlob("pages/*.html")

	// Load the english dictionary
	dict = trieInit()
	dict.readDictionaryFile("english_cleaned.txt")
}

var dataChannels map[string](chan string)

func main() {

	dataChannels = make(map[string](chan string))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uuidString := uuid.New().String()
		cookie := http.Cookie{ Name: "wbClient", Value: uuidString}
		dataChannels[uuidString] = make(chan string)

		http.SetCookie(w, &cookie)
		
		fmt.Printf("index: %s\n",uuidString)
        http.ServeFile(w, r, "pages/index.html")
    })

	http.HandleFunc("/solve", solvePuzzle)

    http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("wbClient")
		uuidString := cookie.Value
		fmt.Printf("start result: %s\n",uuidString)

        // Set headers for SSE
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")

        // Create a context for handling client disconnection
        _, cancel := context.WithCancel(r.Context())
        defer cancel()
				
		var dataChan chan string
		dataChan = dataChannels[uuidString]
		for data := range dataChan {
			fmt.Printf("data: %s\n\n", data)                
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.(http.Flusher).Flush()
		}
		fmt.Printf("end result: %s\n",uuidString)
    })

	http.ListenAndServe(":80", nil)
}

func buildGameBoard(wlInput []string, gbInput []string) (*board, []int) {

	fmt.Println("in buildGameBoard")
	if len(wlInput) == 0 || len(gbInput) == 0 {
		return nil,nil
	}

	gbLowercase := strings.ToLower(gbInput[0])
	gbLowercase = regexp.MustCompile("[^a-z\n]").ReplaceAllString(gbLowercase, "")
	
	wlStrings := strings.Split(strings.ReplaceAll(wlInput[0], " ", ""), ",")
	wlInts := make([]int, len(wlStrings))

	var err error
	wlSum := 0
	for i, wl := range wlStrings {
		wlInts[i], err = strconv.Atoi(wl)
		wlSum += wlInts[i]
		if err != nil {
			return nil,nil
		}
	}

	gbRows := strings.Split(gbLowercase, "\n")
	var gbCleanRows []string
	for _, gbRow := range gbRows {
		if gbRow != "" {
			gbCleanRows = append(gbCleanRows, gbRow)
		}
	}
	gbHeight := len(gbCleanRows)
	if gbHeight * gbHeight != wlSum {
		return nil,nil
	}

	for _, gbRow := range gbCleanRows {
		if len(gbRow) != gbHeight {
			return nil,nil
		}
	}

	return buildBoard(strings.Join(gbCleanRows,""), gbHeight), wlInts
}

func solvePuzzle(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("wbClient")
	uuidString := cookie.Value
	fmt.Printf("start solve: %s\n",uuidString)	
	
	r.ParseForm()
	gb, wl := buildGameBoard(r.Form["word-lengths"], r.Form["game-board"])
	if( gb == nil) {
		fmt.Println("Invalid input")
	}
	gb.printBoard("Game Board: ")
	fmt.Println(wl)

	go func() {

		dataChannels[uuidString] <- "Solving..."
		time.Sleep(1 * time.Second)
		dataChannels[uuidString] <- "Solution!"
	}()

	terr := page.ExecuteTemplate(w, "resultsView.html",nil)
	if terr != nil {
		http.Error(w, terr.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("end solve: %s\n",uuidString)

}
