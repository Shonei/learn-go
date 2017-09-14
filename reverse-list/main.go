package main

import (
	"fmt"
)

type Node struct {
	val  int
	next *Node
}

func main() {
	head := &Node{}
	generateList(head)
	print(head)
	reverse(head)
	// print(head)
}

func generateList(head *Node) {
	temp := head
	for i := 1; i < 15; i++ {
		temp.next = &Node{val: i}
		temp = temp.next
	}
}

func print(head *Node) {

	for temp := head; temp != nil; temp = temp.next {
		fmt.Println(temp.val)
	}
}

func reverse(head *Node) {
	prev := head
	cur := head.next
	next := head.next.next

	prev.next = nil

	for next.next != nil {
		cur.next = prev
		prev = cur
		cur = next
		next = next.next
	}

	cur.next = prev
	next.next = cur
	head = next

	print(head)
}
