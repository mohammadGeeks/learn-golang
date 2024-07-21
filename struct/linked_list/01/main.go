package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) insertAtEnd(d int) {
	// nn is a new node
	nn := &Node{
		data: d,
		next: nil,
	}

	// if our linked list has no nodes, the nn is the linked list head
	if list.head == nil {
		list.head = nn

	} else {
		// loop through, from beginning to end, our linked list
		// the variable 'in' represents the ITERATION NODE in our list
		// when we start the iteration, 'in' is the first node in the list (list.head)
		in := list.head
		// for each iteration, if ITERATION NODE has an 'in.next' pointing to a next node (eg, the pointer is not nil)
		for in.next != nil {
			// set 'in' to the address of the next node in the list
			in = in.next
		}
		// when in.next == nil
		// the loop above will exit
		// this is the node that was the LAST NODE in the list
		// we now have a NEW NODE 'nn' however,
		// so for this last ITERATION NODE that currently does not point to a node beyond it
		// set it to point to the NEW NODE 'nn'
		in.next = nn

		// the NEW NODE 'nn' we created will now be the last node
		// and this NEW NODE will not have a pointer to a next node
		// as this NEW NODE is the new LAST NODE
	}
}

func (list *LinkedList) insertAtPosition(position int, d int) {
	// nn is a new node
	nn := &Node{
		data: d,
		next: nil,
	}

	if position < 1 {
		fmt.Println("Invalid position")
		return
	}

	if position == 1 {
		// if we're inserting into the first position (this is not a zero-based index)
		// then this new node will point to the next node which is the list.head node
		nn.next = list.head
		// the list.head will now be our new node
		list.head = nn
		return
	}

	// 'cn' is current node
	// for inserting at any position other than the first position
	// start with the first node which is our 'list.head'
	cn := list.head
	// then loop through all of the nodes starting at position 2
	// and going all of the way up to the position we want to insert at, less 1
	for i := 2; i <= position-1; i++ {
		// if we have a node here
		if cn != nil {
			// the current node 'cn' then becomes the next node
			cn = cn.next
		}
	}

	// if we have a valid current node 'cn'
	// 'cn' is the node at the position we want to insert at, less 1
	// eg, 'cn' is the node BEFORE the new node 'nn' we want to insert
	if cn != nil {
		// 'cn.next' is pointing to the node after it
		// but as 'nn' will now be before this node
		// we need 'nn.next' to point to the node that 'cn.next' used to point to
		nn.next = cn.next
		// and we need 'cn.next' to point to our new node
		cn.next = nn

		/*
			'an' is another node in this comment

			BEFORE
			cn --> an
			cn pointed to an

			AFTER
			cn --> nn --> an
			nn now points to an
			cn now points to nn
		*/

	} else {
		fmt.Println("Position out of range")
	}
}

func (list *LinkedList) printList() {
	// n is a node
	n := list.head
	for n != nil {
		fmt.Printf("address: %p \tdata: %d \tstructure: %v\n", n, n.data, n)
		n = n.next
	}

}

func main() {

	list := LinkedList{}
	list.insertAtEnd(1)
	list.insertAtEnd(2)
	list.insertAtEnd(3)
	list.insertAtEnd(5)
	list.insertAtPosition(1, 4) // This will insert '4' at position 2
	list.printList()
}
