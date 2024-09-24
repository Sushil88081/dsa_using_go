package main

import "fmt"

type DoublyList struct {
	data int
	prev *DoublyList
	next *DoublyList
}

// CreateList creates a new node and adds it to the end of the list
func CreateList(head *DoublyList, value int) *DoublyList {
	newNode := &DoublyList{data: value, prev: nil, next: nil}

	if head == nil {
		// If the list is empty, return the new node as the head
		return newNode
	}

	// Traverse to the end of the list
	temp := head
	for temp.next != nil {
		temp = temp.next
	}

	// Link the new node at the end
	temp.next = newNode
	newNode.prev = temp

	return head
}

func InsertAtBeg(head *DoublyList, value int) *DoublyList {
	var prev *DoublyList = nil
	var next *DoublyList = nil
	newNode := &DoublyList{data: value, prev: prev, next: next}

	newNode.prev = head
	newNode.next = head
	head = newNode

	return head

}

// PrintList prints the list in both directions
func PrintList(head *DoublyList) {
	fmt.Print("Forward: ")
	for temp := head; temp != nil; temp = temp.next {
		fmt.Printf("%d <-> ", temp.data)
	}
	fmt.Println("nil")

}

func main() {
	var head *DoublyList
	fmt.Println("Enter values to insert into your list:")
	var value int
	for i := 0; i < 5; i++ { // Change loop to run 5 times
		fmt.Scan(&value)
		head = CreateList(head, value)
	}
	PrintList(head)
	fmt.Println("Enter values which you insert at begining of list:")
	fmt.Scan(&value)
	head = InsertAtBeg(head, value)
	PrintList(head)
}
