package main


type trieNode struct {
    children [26]*trieNode
    wordEnds bool
}

type trie struct {
    root *trieNode
}

//inititlaizing a new trie 
func trieInit() *trie {
    t := new(trie)
    t.root = new(trieNode)
    return t
}

// Inserting words into trie
func (t *trie) insert(word string) {
    current := t.root
    for _, letter := range word {
        index := letter - 'a'

        if current.children[index] == nil {
            current.children[index] = new(trieNode)
        }
        current = current.children[index]
    }
    current.wordEnds = true
}

// Searching for words
// 1  = complete word found
// -1 = trie branch exists, but it is an incomplete word
// 0  = no such branch found
func (t *trie) search(word string) int {
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

// // Search branch beginning with a specific trie node
// func (node *trieNode) searchPartial(letter rune) *trieNode {
//     index := letter - 'a'
//     return node.children[index]
// }

