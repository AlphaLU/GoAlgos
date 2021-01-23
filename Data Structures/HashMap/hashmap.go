package main

import (
	"fmt"
)

var p = fmt.Println
const tableLen = 7

type HashTable struct {
	array [tableLen]*LinkedList
}

type LinkedList struct {
	head *Node
}

type Node struct {
	key string
	next *Node
}

//Hash table methods
func Init() *HashTable {
	table := &HashTable{}
	for i := range table.array {
		table.array[i] = &LinkedList{}
	}
	return table
}

//Insert
func (h *HashTable) Insert(key string) *Node {
	index := hash(key)
	return h.array[index].Insert(key)
}

//Search
func (h *HashTable) Search(key string) {
	index := hash(key)
	h.array[index].Search(key)
}

//Delete
func (h *HashTable) Delete(key string) {
	index := hash(key)
	p(h.array[index].Delete(key))
}

//Get all entries
func (h *HashTable) GetAllEntries() {
	for i := 0; i < tableLen; i++ {
		p("Entries under index:", i)
		h.array[i].ListItems()
	}
}

//Linked list methods
//Insert
func (l *LinkedList) Insert(key string) *Node {
	node := &Node{key: key}
	node.next = l.head
	l.head = node
	return node
}

//Search
func (l *LinkedList) Search(key string) {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.key == key {
			p("Found,", key, "In node", currentNode)
			break
		} else {
			if currentNode.next != nil {
				currentNode = currentNode.next
			} else {
				p(key, "Not Found")
				break
			}
		}
	}
}

//Delete
func (l *LinkedList) Delete(key string) bool {
	currentNode := l.head
	for currentNode != nil {
		if currentNode.key == key {
			if currentNode.next != nil {
				if currentNode == l.head {
					temp := currentNode.next.next
					l.head = currentNode.next
					l.head.next = temp
					return true
				} else {
					currentNode = currentNode.next
				}
			} else {
				if currentNode == l.head {
					l.head = nil
				} else { 
					currentNode = currentNode.next
				}
			}
		} else {
			currentNode = currentNode.next
		}
	}
	return false
}

//List items in the array
func (l *LinkedList) ListItems() {
	currentNode := l.head
	if currentNode == nil {
		p("None")
	} else {
		for currentNode != nil {
			p("Value:", currentNode.key)
			if currentNode.next != nil {
				currentNode = currentNode.next
			} else {
				break
			}
		}
	}
}

//Hash function, can improve
func hash(key string) int {
	index := 0
	for _, val := range key {
		index += int(val)
	}
	return index % tableLen
}

func main() {
	hashTable := Init()
	p(hash("Liran"))
	p(hash("rassaalf"))
	p(hashTable)
	hashTable.Insert("Liran")
	hashTable.Insert("Test")
	hashTable.Insert("Username")
	hashTable.Insert("Alert")
	hashTable.Insert("GoLang")
	hashTable.Insert("rassaalf")
	hashTable.Insert("rassaalf")
	hashTable.Insert("rassaalf")

	p(hashTable.array[hash("rassaalf")].head.next.next.next)

	
	hashTable.Search("rassaalf")
	// hashTable.Search("Liran")
	// hashTable.Search("Test")
	// hashTable.Search("Username")
	// hashTable.Search("Alert")
	// hashTable.Search("GoLang")

	hashTable.Delete("rassaalf")

	hashTable.Search("rassaalf")

	hashTable.GetAllEntries()
	hashTable.Delete("Username")
	hashTable.GetAllEntries()
	hashTable.Delete("Test")
	hashTable.GetAllEntries()
}