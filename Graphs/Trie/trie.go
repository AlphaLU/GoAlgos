package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var p = fmt.Println
const alphabet = 26

type Node struct {
	children [alphabet]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	newTrie := &Trie{root: &Node{}}
	return newTrie
}

func (t *Trie) InsertWord(w string) {
	w = strings.ToLower(w)
	wLen := len(w)
	currentNode := t.root
	for i := 0; i < wLen; i++ {
		charIndex := w[i] - 'a'
		if (currentNode.children[charIndex] == nil) {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) SearchWord(w string) (bool, *Node) {
	w = strings.ToLower(w)
	wLen := len(w)
	checkIndex := 0
	currentNode := t.root
	for i := 0; i < wLen; i++ {
		charIndex := w[i] - 'a'
		if (currentNode.children[charIndex] != nil) {
			checkIndex++
			currentNode = currentNode.children[charIndex]
		} else {
			return false, nil
		}
	}
	if currentNode.isEnd == true {
		return true, currentNode
	}

	return false, nil
}

func (root *Node) CountWords() int {
	wordCount := 0

	if root.isEnd {
		wordCount++
	}
	for i := 0; i < alphabet; i++ {
		if (root.children[i] != nil) {
			wordCount += root.children[i].CountWords()
		}
	}
	return wordCount
}

func main() {

	test := InitTrie()
	startTime := time.Now()

	res := InsertMany(test)
	endTime := time.Since(startTime)
	p("Time took to insert: ", endTime.Seconds())

	if (res) {
		var searches []string = []string{"Liran", "Smaugg", "Gandalf", "TESTING"}
		for i := 0; i < len(searches); i++ {
			isExist, _ := test.SearchWord(searches[i])
			
			if isExist {
				p(searches[i], " Exists: ", isExist)
			} else {
				p(searches[i], " Exists: ", isExist)
			}
		}
	}
	p("Counting words...")
	startTime = time.Now()
	p("There are: ", test.root.CountWords(), " words")
	endTime = time.Since(startTime)
	p("Time took to count words: ", endTime.Seconds())
	
}


//Driver
func InsertMany(t *Trie) bool {
	f, err := os.Open("words.in")
	if err != nil {
		p("Can't open file")
		return false
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	var word string
	for {
		line, err := reader.ReadString('\n')
		if err != io.EOF && err != nil {
			break
		}
		word = string(line)
		word := strings.Split(word, " ")
		for j := 0; j < len(word); j++ {
			if (len(word[0]) > 0) {
			t.InsertWord(word[j])
			}
		}
		if err != nil {
			break
		}
	}
	return true
}