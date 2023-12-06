package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"slices"
	"time"
	"context"
	"github.com/google/uuid"
	// "strings"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uuidString := uuid.New().String()
		cookie := http.Cookie{ Name: "wbClient", Value: uuidString}
		dataChannels[uuidString] = make(chan string)

		http.SetCookie(w, &cookie)
		
		fmt.Printf("index: %s\n",uuidString)
        http.ServeFile(w, r, "pages/index.html")
    })

	http.HandleFunc("/solve", solvePuzzle)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	// http.HandleFunc("/static/", func(wr http.ResponseWriter,req *http.Request) {
	// 	// Determine mime type based on the URL
	// 	if strings.HasSuffix(req.URL.Path, ".css") {
	// 		wr.Header().Set("Content-Type","text/css")
	// 	} else if strings.HasSuffix(req.URL.Path, ".js") {
	// 		wr.Header().Set("Content-Type","application/javascript")
	// 	} 
	// 	http.StripPrefix("/static/", fs).ServeHTTP(wr,req)
	// })

    http.HandleFunc("/result", func(w http.ResponseWriter, r *http.Request) {
		cookie, _ := r.Cookie("wbClient")
		uuidString := cookie.Value
		fmt.Printf("start result: %s\n",uuidString)

        // Set headers for SSE
        w.Header().Set("Content-Type", "text/event-stream")
        w.Header().Set("Cache-Control", "no-cache")
        w.Header().Set("Connection", "keep-alive")

        // Create a channel to send data
        // dataCh := make(chan string)

        // Create a context for handling client disconnection
        _, cancel := context.WithCancel(r.Context())
        defer cancel()

        // Send data to the client
		
		var dataChan chan string
		dataChan = dataChannels[uuidString]
        // go func() {
		// data := <-dataChan
		for data := range dataChan {
			fmt.Printf("data: %s\n\n", data)                
			fmt.Fprintf(w, "data: %s\n\n", data)
			w.(http.Flusher).Flush()

			// data = <-dataChan
			// fmt.Printf("data: %s\n\n", data)                
			// fmt.Fprintf(w, "data: %s\n\n", data)
			// w.(http.Flusher).Flush()
		}
        // }()

		// Simulate sending data periodically
		// for {
		// 	time.Sleep(1 * time.Second)
		// 	dataChan <- time.Now().Format(time.TimeOnly)
			
		// }
		fmt.Printf("end result: %s\n",uuidString)


    })

	http.ListenAndServe(":80", nil)
}


type Thing struct {
	Text   string
	Things []Thing
}

func solvePuzzle(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("wbClient")
	uuidString := cookie.Value
	fmt.Printf("start solve: %s\n",uuidString)

	
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

func oldTestMain() {
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
	findPhrase(0, wordLengths, &rtn, dict)
	// printSubtree(&rtn, "",len(wordLengths))

	flattened := make(map[string]int,0)
	collapseSubtree(&rtn, "", len(wordLengths), &flattened)
	answers := make([]string, len(flattened))
	i := 0
	for k,_ := range flattened {
		answers[i] = k
		i++
	}
	
	slices.Sort(answers)
	
	f, _ := os.Create("out.txt")


	for _,a := range answers {
		fmt.Println(a)
		f.WriteString(fmt.Sprintf("%s\n", a))
	}
	f.Close()
}