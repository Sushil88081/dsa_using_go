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

func DeleteAtFirst(head *Node) *Node {
	var temp *Node
	temp = head
	head = head.next
	temp.next = nil
	fmt.Println("First node deleted")
	PrintList(head)
	return head
}

func DeleteOnlyNode(head *Node) *Node {
	head = nil
	fmt.Println("Only node deleted")
	PrintList(head)
	return head
}

func DeleteAtPosition(head *Node, position int) *Node {
	// If linked list is empty
	if head == nil {
		return nil
	}

	// If the position is 0, delete the head node
	if position == 0 {
		return head.next
	}

	var temp *Node = head

	// Traverse the list to find the node before the position
	for i := 0; temp != nil && i < position-1; i++ {
		temp = temp.next
	}

	// If position is greater than the number of nodes
	if temp == nil || temp.next == nil {
		return head
	}

	// Remove the node at the given position
	temp.next = temp.next.next

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
	// DeleteAtFirst(head)
	// DeleteOnlyNode(head)
	DeleteAtPosition(head, 2)
	PrintList(head)

}
