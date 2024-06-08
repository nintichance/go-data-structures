package main

import (
	"fmt"
	"strings"
)

func main() {

	dt, dts := getDataType()
	for !searchSlice(dt, dts) {
		fmt.Printf("Your input: %v is invalid. Please try again:\n", dt)
		fmt.Scan(&dt)
	}

	fmt.Printf("Wonderful! You chose: %v\n", dt)

	if dt == "linkedlist" {
		runLinkedList()
	}
}

func getDataType() (string, []string) {
	var dataType string
	var dataTypes []string
	dsString := "linkedlist stack queue bst minheap hash trie"
	dataTypes = strings.Fields(dsString)

	fmt.Println("welcome to ninti's data structures in go!")
	fmt.Println("which data structure would you like to see?")
	fmt.Println("choose one: linkedlist, stack, queue, bst, minheap, hash, trie")
	fmt.Scan(&dataType)
	return dataType, dataTypes
}
func searchSlice(str string, arr []string) bool {
	var found bool = false
	for i := 0; i < len(arr); i++ {
		if arr[i] == str {
			found = true
			break
		}
	}
	return found
}

func runLinkedList() {
	l := &LinkedList{}
	l.init(1)
	l.insert(2)
	l.insert(3)
	l.insert(4)
	fmt.Printf("The size is: %v\n", l.size())
	l.print(l.Head)
	fmt.Println("##############")
	l.printReverse(l.Head)
	fmt.Println("REVERSE ##############")
	l.reverse()
	l.print(l.Head)
	fmt.Println("REVERSE RECURSIVE ##############")
	l.reverseRecursive(l.Head)
	l.print(l.Head)
	fmt.Println("##############")
}
