package main

import (
    "os"
    "log"
    "bufio"
)

type dictionaryTrieNode struct {
    children [26]*dictionaryTrieNode
    wordEnds bool
}

// Search branch beginning with a specific trie node
func (n *dictionaryTrieNode) findChildNode(letter rune) *dictionaryTrieNode {
    index := letter - 'a'
    return n.children[index]
}

type dictionaryTrie struct {
    root *dictionaryTrieNode
}

//inititlaizing a new dictionary trie 
func trieInit() *dictionaryTrie {
    t := new(dictionaryTrie)
    t.root = new(dictionaryTrieNode)
    return t
}

// Inserting words into trie
func (t *dictionaryTrie) insert(word string) {
    current := t.root
    for _, letter := range word {
        index := letter - 'a'

        if current.children[index] == nil {
            current.children[index] = new(dictionaryTrieNode)
        }
        current = current.children[index]
    }
    current.wordEnds = true
}

// Searching for words
// 1  = complete word found
// -1 = trie branch exists, but it is an incomplete word
// 0  = no such branch found
func (t *dictionaryTrie) search(word string) int {
    node := t.root
    for _, letter := range word {
        index := letter - 'a'
        if node.children[index] == nil {
            return 0
        }
        node = node.children[index]
    }

    if node.wordEnds {
        return 1
    }
    return -1
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



