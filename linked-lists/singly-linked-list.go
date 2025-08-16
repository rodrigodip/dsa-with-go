package linkedLists

import (
	"fmt"
)

type LinkedList struct {
	head *Node
}
type Node struct {
	value string
	next  *Node
}

// PERF: Operations:
// Insert at beginning
// Insert at end
// Delete first
// Delete last
// Delete by value
// Print list
// Get length

// TODO: Operations
// Insert by position
// Delete by Position
// Search value

// Insert a element at the beginning of the list
func (ll *LinkedList) InsertAtBeginning(data string) {
	newNode := &Node{value: data, next: ll.head}
	ll.head = newNode
}

// Insert a element at the end of the list
func (ll *LinkedList) InsertAtEnd(data string) {
	newNode := &Node{value: data, next: nil}

	if ll.head == nil {
		ll.head = newNode
		return
	}

	current := ll.head

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

// Deletes the first element on the list
func (ll *LinkedList) DeleteFirst() {
	if ll.head == nil || ll.head.next == nil {
		ll.head = nil
		return
	}

	ll.head = ll.head.next
}

// Deletes the last element on the list
func (ll *LinkedList) DeleteLast() {
	if ll.head == nil || ll.head.next == nil {
		ll.head = nil
		return
	}

	current := ll.head

	for current.next.next != nil {
		current = current.next
	}

	current.next = nil
}

// Deletes a element by value
func (ll *LinkedList) DeleteValue(value string) {
	if ll.head == nil {
		return
	}

	if ll.head.value == value {
		ll.head = ll.head.next
		return
	}

	current := ll.head

	for current.next != nil {
		if current.next.value == value {
			current.next = current.next.next
			return
		}
		current = current.next
	}
}

// Print all elements on the list
func (ll *LinkedList) PrintList() {
	if ll.head == nil {
		fmt.Println("-- Empty List --")
		return
	}

	current := ll.head

	for current.next != nil {
		fmt.Print(current.value, " --> ")
		current = current.next
	}

	fmt.Println(current.value, "--> Null")
}

// Returns the number of elements on the list
func (ll *LinkedList) GetLen() uint {
	if ll.head == nil {
		return 0
	}

	current := ll.head
	var count uint

	for current.next != nil {
		count++
		current = current.next
	}

	count++
	return count
}
