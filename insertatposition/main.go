package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func createNode(data int) *Node {
	newNode := &Node{data: data, next: nil}
	return newNode
}
func printList(head *Node) {

	p := head
	for p != nil {
		fmt.Print(p.data, " ->")
		p = p.next
	}
	fmt.Println("nil")

}

func insertAtPosition(data int, pos int, head *Node) *Node {
	newNode := createNode(data)
	p := head
	// If the position is 1 (or 0-based index 0), insert at the head
	if pos == 1 {
		newNode.next = head
		return newNode // New head
	}

	for i := 0; i < pos-1 && p != nil; i++ {
		p = p.next
	}

	// If the position is out of bounds
	if p == nil {
		fmt.Println("Position out of bounds")
		return head
	}

	newNode.next = p.next
	p.next = newNode
	return head
}

func main() {
	var head *Node
	head = nil
	var p *Node
	var data int
	var pos int

	fmt.Println("Enter number of link list")
	for i := 0; i < 5; i++ {
		fmt.Scan(&data)
		newNode := createNode(data)
		if head == nil {
			head = newNode
			p = newNode

		} else {
			p.next = newNode
			p = newNode
		}

	}
	printList(head)
	fmt.Print("Enter data to insert: ")
	fmt.Scan(&data)
	fmt.Println("Enter position to insert in link list")
	fmt.Scan(&pos)
	head = insertAtPosition(data, pos, head)
	fmt.Print("Updated linked list: ")
	printList(head)

}
