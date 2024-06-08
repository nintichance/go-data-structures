package main

import "fmt"

func main() {
	var dataType string
	fmt.Println("welcome to ninti's data structures in go!")
	fmt.Println("which data structure would you like to see?")
	fmt.Println("choose one: linked list, stack, queue, bst, priority queue, hash, trie")
	fmt.Scan(&dataType)
	fmt.Printf("Wonderful! You chose: %v", dataType)
}
