package main

import (
	"net/http"
	"fmt"
	"github.com/google/uuid"
	"context"
	"strings"
	"regexp"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	uuidString := uuid.New().String()
	cookie := http.Cookie{ Name: "wbClient", Value: uuidString}
	commsChannels[uuidString] = make(chan string)
	dataChannels[uuidString] = make(chan resultTreeNode)

	http.SetCookie(w, &cookie)
	
	fmt.Printf("index: %s\n",uuidString)
	http.ServeFile(w, r, "pages/index.html")
}

func handlePuzzleInput(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("wbClient")
	uuidString := cookie.Value
	fmt.Printf("start solve: %s\n",uuidString)	
	commsChannel := commsChannels[uuidString]
	r.ParseForm()

	go func() {
		commsChannel <- "Solving."

		
		gb, wl, err := initGameFromUserInput(r.Form["word-lengths"], r.Form["game-board"])
		if(err != nil) {
			commsChannel <- fmt.Sprintf("Error: %s", err)
			return
		}

		gb.printBoard("Game Board: ")
		fmt.Println(wl)

		commsChannel <- "Solving.."
		resultRoot := resultTreeNode {
			Word: "-- ROOT --",
			gridIndices: []int{},
			collapsedBoard: gb,
			sourceBoard: gb,
			NextWords: nil,
		}
		commsChannel <- "Solving..."
		findPhrase(0, wl, &resultRoot, dict)

		
		commsChannel <- "Solution FOUND!"
		
		commsChannel <- "Solution!"
		dataChannels[uuidString] <- resultRoot
	}()

	terr := page.ExecuteTemplate(w, "resultsView.html",nil)
	if terr != nil {
		http.Error(w, terr.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("end solve: %s\n",uuidString)
}

func initGameFromUserInput(wlInput []string, gbInput []string) (*board, []int, error) {

	fmt.Println("in buildGameBoard")
	if len(wlInput) == 0 || len(gbInput) == 0 {
		return nil,nil, fmt.Errorf("missing an input")
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
			return nil,nil, fmt.Errorf("invalid word length number: \"%s\"", wl)
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
		return nil,nil, fmt.Errorf("sum of word lengths does not match game board size")
	}

	for ri, gbRow := range gbCleanRows {
		if len(gbRow) != gbHeight {
			return nil,nil, fmt.Errorf("game board is not square on row %d", ri+1)
		}
	}

	return buildBoard(strings.Join(gbCleanRows,""), gbHeight), wlInts, nil
}

func sendResultOverSSE(w http.ResponseWriter, r *http.Request) {
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
			
	var commsChan chan string
	var result resultTreeNode
	commsChan = commsChannels[uuidString]
	for data := range commsChan {
		if data == "Solution!" {


			result = <- dataChannels[uuidString]
			if result.NextWords != nil { 
				builder := &strings.Builder{}
				if err := page.ExecuteTemplate(builder, "resultsTree.html", result.NextWords); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				str := builder.String()
				str = strings.ReplaceAll(str, "\r\n", "")
				fmt.Fprintf(w, "data: %s\n\n", str)
			} else {
				fmt.Printf("data: Empty Solution%s\n\n", data)   
				fmt.Fprintf(w, "data: Empty Solution%s\n\n", data)
			}
		} else {
			fmt.Printf("data: %s\n\n", data)   
			fmt.Fprintf(w, "data: %s\n\n", data)
		}
					 
		w.(http.Flusher).Flush()
	}
	fmt.Printf("end result: %s\n",uuidString)
}

