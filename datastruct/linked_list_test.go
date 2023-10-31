package datastruct

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Create list
	l := &List{}
	fmt.Println("Initial Length", l.Length())
	// Append items
	l.Append(3)
	l.Append(2)
	l.Append(1)
	l.Append(10)
	l.Print()
	// Prepend items
	l.Prepend(20)
	l.Print()
	// Length
	fmt.Println("Length", l.Length())
	// Contains
	fmt.Println("Contains 100", l.Contains(100))
	fmt.Println("Contains 20", l.Contains(1))
	// Index
	fmt.Println("Position of 20", l.Index(20))
	fmt.Println("Position of 100", l.Index(100))
	// Delete first
	l.DeleteFirst()
	l.DeleteFirst()
	l.Print()

	// Delete data
	l.Delete(300)
	l.Print()
	l.Delete(2)
	l.Print()
}
