package main

import "fmt"

type Node struct {
	Val  uint
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) init(val uint32) {
	l.Head = &Node{
		Val: uint(val),
	}
}

func (l *LinkedList) size() uint {
	var count uint
	node := l.Head

	for node != nil {
		count += 1
		node = node.Next
	}
	return count
}

func (l *LinkedList) insert(val uint32) {
	newNode := Node{
		Val: uint(val),
	}

	head := l.Head

	for head.Next != nil {
		head = head.Next
	}
	head.Next = &newNode
}

func (l *LinkedList) print(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("PRINTED %v\n", n.Val)
	l.print(n.Next)
}

func (l *LinkedList) printReverse(n *Node) {
	if n == nil {
		return
	}
	l.printReverse(n.Next)
	fmt.Printf("PRINTED %v\n", n.Val)
}

func (l *LinkedList) reverse() {
	curr := l.Head
	var prev *Node
	var next *Node

	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}
