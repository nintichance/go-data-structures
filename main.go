package main

import (
	"fmt"
	"strings"
)

func main() {
	var dataType string
	var dataTypes []string
	dsString := "linkedlist stack queue bst minheap hash trie"
	dataTypes = strings.Fields(dsString)

	fmt.Println("welcome to ninti's data structures in go!")
	fmt.Println("which data structure would you like to see?")
	fmt.Println("choose one: linkedlist, stack, queue, bst, minheap, hash, trie")
	fmt.Scan(&dataType)

	for !searchSlice(dataType, dataTypes) {
		fmt.Printf("Your input: %v is invalid. Please try again:", dataType)
		fmt.Scan(&dataType)
	}

	fmt.Printf("Wonderful! You chose: %v", dataType)
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
