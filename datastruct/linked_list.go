package datastruct

import "fmt"

// Linked list in Go

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
}

// Append an item to the end of the list
func (l *List) Append(number int) {
	// Create the new node
	node := &Node{Data: number, Next: nil}

	// 2 cases
	if l.Head == nil {
		// When list empty
		l.Head = node
	} else {
		// When list has values
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
	}
}

func (l List) Print() {
	fmt.Print("Head")

	// Short version
	for p := l.Head; p != nil; p = p.Next {
		fmt.Print(" <- ", p.Data)
	}
	fmt.Println()

	// Long version
	// p := l.Head
	// for p != nil {
	//	p = p.Next
	//	fmt.Print(" <- ", p.Data)
	// }
}

func (l *List) Prepend(number int) {
	// Create the node
	node := &Node{Data: number, Next: nil}

	// 2 cases
	// When list is empty
	if l.Head == nil {
		l.Head = node
	} else {
		temp := l.Head
		node.Next = temp
		l.Head = node
	}
}

func (l List) Length() int {
	p := l.Head
	var len int
	for p != nil {
		len++
		p = p.Next
	}
	return len
}

func (l List) Contains(number int) bool {
	p := l.Head
	for p != nil {
		if p.Data == number {
			return true
		}
		p = p.Next
	}
	return false
}

func (l List) Index(number int) int {
	p := l.Head
	var index int
	var isFound bool
	for p != nil {
		if p.Data == number {
			isFound = true
			break
		}
		index++
		p = p.Next
	}
	if isFound {
		return index
	}
	return -1
}

func (l *List) DeleteFirst() {
	if l.Head != nil {
		l.Head = l.Head.Next
	}
}

func (l *List) Delete(data int) {
	if l.Head != nil {
		if l.Head.Data == data {
			l.Head = l.Head.Next
		} else {
			curr := l.Head
			prev := &Node{}
			var isFound bool
			for curr.Next != nil {
				if curr.Data == data {
					isFound = true
					break
				}
				prev = curr
				curr = curr.Next
			}
			if isFound {
				prev.Next = curr.Next
			}
		}
	}
}
