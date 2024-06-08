package main

import (
	"fmt"
	"strings"
)

func main() {

	dt, dts := getDataType()
	for !searchSlice(dt, dts) {
		fmt.Printf("your input: %v is invalid. please try again:\n", dt)
		fmt.Scan(&dt)
	}

	fmt.Printf("wonderful! you chose: %v\n", dt)

	switch dt {
	case "linkedlist":
		runLinkedList()
	case "stack":
		var str string

		fmt.Println("enter a string of characters: () {} [] in any combination--balanced or unbalanced")
		fmt.Scan(&str)

		for !isValidString(str) {
			fmt.Printf("your input: %v is invalid. \nenter a string of characters: () {} [] in any combination--balanced or unbalanced\n", str)
			fmt.Scan(&str)
		}

		if balanceParens(str) {
			fmt.Printf("The input: %v is balanced\n", str)
		} else {
			fmt.Printf("The input: %v is not balanced\n", str)
		}
	default:
		return
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
	fmt.Println("##############")
}

func isValidString(str string) bool {
	isValid := false
	for i := 0; i < len(str); i++ {
		fmt.Println("THE STRING", str[i])

		if string(str[i]) == "{" ||
			string(str[i]) == "}" ||
			string(str[i]) == "(" ||
			string(str[i]) == ")" ||
			string(str[i]) == "[" ||
			string(str[i]) == "]" {
			isValid = true
		}
	}

	return isValid
}
