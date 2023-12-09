package main

import (
	"html/template"
	"net/http"
)

var page *template.Template
var dict *dictionaryTrie

func init() {
	page, _ = template.ParseGlob("pages/*.html")

	// Load the english dictionary
	dict = trieInit()
	dict.readDictionaryFile("english_cleaned.txt")
	
	// make channels to communicate between functions
	commsChannels = make(map[string](chan string), 1)
	dataChannels = make(map[string](chan resultTreeNode), 1)
}

var commsChannels map[string](chan string)
var dataChannels map[string](chan resultTreeNode)

func main() {

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", index)

	http.HandleFunc("/solve", handlePuzzleInput)

    http.HandleFunc("/result", sendResultOverSSE)

	http.ListenAndServe(":80", nil)
}




