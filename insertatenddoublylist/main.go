package main

import (
	"fmt"
)

type DoublyList struct {
	data int
	prev *DoublyList
	next *DoublyList // Pointer to next node in DLL
}

func CreateList(head *DoublyList) *DoublyList {

	var value int
	fmt.Println("Enter items in the list:")

	// Input 5 items
	for i := 0; i < 5; i++ {
		fmt.Scan(&value)
		newNode := &DoublyList{data: value}

		if head == nil {
			// If the list is empty, newNode becomes the head
			head = newNode
		} else {
			// Traverse to the end of the list
			temp := head
			for temp.next != nil {
				temp = temp.next
			}
			// Link the new node to the end of the list
			temp.next = newNode
			newNode.prev = temp
		}
	}
	return head
}

func InsertAtEnd(head *DoublyList) *DoublyList {

	var value int
	fmt.Println("Enter insert items in the list:")
	fmt.Scan(&value)
	var newNode *DoublyList = &DoublyList{data: value}
	// If the list is empty, newNode becomes the head
	if head == nil {
		return newNode
	}
	var ptr *DoublyList
	ptr = head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = newNode
	newNode.prev = ptr

	return head
}
func PrintList(head *DoublyList) {
	fmt.Print("Forward: ")
	for temp := head; temp != nil; temp = temp.next {
		fmt.Printf("%d <-> ", temp.data)
	}
	fmt.Println("nil")

}

func main() {
	var head *DoublyList = nil
	head = CreateList(head)
	PrintList(head)
	head = InsertAtEnd(head)
	PrintList(head)
}
