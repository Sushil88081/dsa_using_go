package main

import "fmt"

type Node struct {
	data int
	next *Node
}

// Function to create a new node
func createNode(data int) *Node {
	newNode := &Node{data: data, next: nil}
	return newNode
}

func ReverseList(head *Node) *Node {
	var current = head
	var prev *Node = nil
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	head = prev
	PrintList(head)
	return head
}
func delete(head *Node) *Node {
	var ptr *Node
	ptr = head
	head = ptr.next
	ptr.next = head.next

	return head
}

// Function to print the linked list
func PrintList(head *Node) {
	p := head // Initialize p to point to head of the list
	for p != nil {
		fmt.Printf("%d -> ", p.data)
		p = p.next
	}
	fmt.Println("nil") // Indicate the end of the list
}

func main() {
	fmt.Println("Linked List creation")

	var head *Node = nil
	var p *Node
	var data int

	fmt.Println("Enter numbers in the list:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&data)
		newNode := createNode(data)
		if head == nil {
			head = newNode
			p = head
		} else {
			p.next = newNode
			p = newNode
		}
	}

	fmt.Print("The linked list is: ")
	PrintList(head)
	// ReverseList(head)
	head = delete(head)
	PrintList(head)

}
