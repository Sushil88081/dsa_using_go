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

func deleteFirstNode(head *DoublyList) *DoublyList {

	head = head.next
	head.prev = nil
	return head
}

func deleteBetweenNodes(head *DoublyList, position int) *DoublyList {
	// If the list is empty, return nil
	if head == nil {
		return nil
	}

	temp := head
	// Traverse to the node just before the desired position
	for i := 1; i < position && temp != nil; i++ {
		temp = temp.next
	}

	// If the position is greater than the length of the list
	if temp == nil {
		return head // Nothing to delete
	}

	// At this point, temp is the node at the given position
	// Update pointers to remove temp from the list
	if temp.prev != nil {
		temp.prev.next = temp.next // Link the previous node to the next node
	} else {
		// If temp is the head, move head to the next node
		head = temp.next
	}

	if temp.next != nil {
		temp.next.prev = temp.prev // Link the next node back to the previous node
	}

	// Optional: Free the node if necessary (in languages with manual memory management)
	return head // Return the (possibly new) head of the list
}

func deleteLastNode(head *DoublyList) *DoublyList {
	temp := head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = nil
	temp.prev.next = nil
	return head
}

// PrintList prints the list in both directions
func PrintList(head *DoublyList) {
	fmt.Print("Forward: ")
	for temp := head; temp != nil; temp = temp.next {
		fmt.Printf("%d <-> ", temp.data)
	}
	fmt.Println("nil")

	// // To print in reverse
	// fmt.Print("Backward: ")
	// temp := head
	// // Traverse to the end
	// for temp.next != nil {
	// 	temp = temp.next
	// }
	// // Print backwards
	// for temp != nil {
	// 	fmt.Printf("%d <-> ", temp.data)
	// 	temp = temp.prev
	// }
	// fmt.Println("nil")
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
	fmt.Println("Deleting  node...")
	var position int
	fmt.Println("Enter position to delete")
	fmt.Scan(&position)
	head = deleteBetweenNodes(head, position)
	// head = deleteLastNode(head)
	PrintList(head)

}
