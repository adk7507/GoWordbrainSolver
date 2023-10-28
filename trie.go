package main


type dictionaryNode struct {
    children [26]*dictionaryNode
    wordEnds bool
}

// Search branch beginning with a specific trie node
func (n *dictionaryNode) findChildNode(letter rune) *dictionaryNode {
    index := letter - 'a'
    return n.children[index]
}

type dictionaryTrie struct {
    root *dictionaryNode
}

//inititlaizing a new dictionary trie 
func trieInit() *dictionaryTrie {
    t := new(dictionaryTrie)
    t.root = new(dictionaryNode)
    return t
}

// Inserting words into trie
func (t *dictionaryTrie) insert(word string) {
    current := t.root
    for _, letter := range word {
        index := letter - 'a'

        if current.children[index] == nil {
            current.children[index] = new(dictionaryNode)
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



