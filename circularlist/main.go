package main

import "fmt"

type CircularList struct {
	data int
	next *CircularList
}

// CreateCircularList inserts a new node at the end of the circular list
func CreateCircularList(head *CircularList, value int) *CircularList {
	newNode := &CircularList{data: value}

	// If the list is empty, initialize the circular structure
	if head == nil {
		newNode.next = newNode // Point the new node to itself to make it circular
		return newNode
	}

	// Traverse to the last node
	ptr := head
	for ptr.next != head {
		ptr = ptr.next
	}

	// Insert the new node at the end and make it circular
	ptr.next = newNode
	newNode.next = head

	return head
}

// PrintCircularList prints the circular linked list
func PrintCircularList(head *CircularList) {
	if head == nil {
		fmt.Println("List is empty")
		return
	}

	ptr := head
	for {
		fmt.Printf("%d -> ", ptr.data)
		ptr = ptr.next
		if ptr == head { // When we reach the head again, stop to avoid infinite loop
			break
		}
	}
	fmt.Println("(circular)")
}

func main() {
	var head *CircularList
	var n, value int

	fmt.Println("Enter the number of nodes you want to create in the circular list:")
	fmt.Scan(&n)

	// Create the circular list by inserting 'n' nodes
	for i := 0; i < n; i++ {
		fmt.Println("Enter value for node", i+1)
		fmt.Scan(&value)
		head = CreateCircularList(head, value)
	}

	// Print the circular list
	PrintCircularList(head)
}
