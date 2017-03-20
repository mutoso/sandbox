// © 2017 Alastair Feille
// Licensed under the MIT License
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Comparable interface {
	Compare(b interface{}) int
}

type MyInt int

func (a MyInt) Compare(b interface{}) int {
	if int(a) < int(b.(MyInt)) {
		return -1
	} else if int(a) > int(b.(MyInt)) {
		return 1
	} else {
		return 0
	}
}

type Node struct {
	name string
}

func (a Node) Compare(b interface{}) int {
	if a.name < b.(Node).name {
		return -1
	} else if a.name > b.(Node).name {
		return 1
	} else {
		return 0
	}
}
func (a Node) String() string {
	return a.name
}

func main() {
	numbers := make([]Comparable, 20)

	// fill array with random numbers
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(numbers); i++ {
		numbers[i] = MyInt(rand.Intn(100))
	}

	fmt.Println("Original:", numbers)
	heapsort(numbers)
	fmt.Println("Sorted:", numbers)
	fmt.Println()

	// Names
	names := []string{"Brandee", "Nelida", "Jaqueline", "Candyce", "Wayne", "Anissa", "Randal", "Milton", "Manda", "Pasquale", "Alpha", "Destiny", "Romaine", "Waneta", "Claudio", "Arnulfo", "Yukiko", "Barbra", "Judie", "Larry"}

	nodes := make([]Comparable, len(names))
	for i, d := range names {
		nodes[i] = Node{name: d}
	}

	fmt.Println("Original:", nodes)
	heapsort(nodes)
	fmt.Println("Sorted:", nodes)
}

func heapsort(array []Comparable) {
	// Step 1: Heapify
	// rearrange the array into heap
	// by going from the right to left, bottom to top
	// and moving each node down if it violates the
	// heap order property (that is, if it is larger
	// than its parent, since this is a *max* heap)
	for i := len(array)/2 - 1; i >= 0; i-- {
		siftDown(i, array, len(array))
	}

	// Step 2: Sort
	// from the last element to the first
	for i := len(array) - 1; i >= 0; i-- {
		// pull out the largest element from the heap root
		// this is the same as a deleteMax() call
		v := array[0]
		// set the new root to the last element in the heap
		array[0] = array[i]
		// move that new root down into its proper place
		siftDown(0, array, i)
		// the heap will shrink as i decreases
		// so put the largest element that we pulled out of
		// the heap right after the end of our shrinking heap
		// at index i
		array[i] = v
	}
	// the heap size is now zero
	// and all elements are sorted in the array
}

// siftDown only actually runs on half of the array,
// so it actually runs in O(n) instead of O(nlogn)
func siftDown(i int, array []Comparable, n int) {
	// parent(i) = (i-1)/2
	// leftchild(i) = 2i+1
	// rightchild(i) = 2i+2.

	// we want to move temp down the heap
	// until it fulfills the heap-order property
	temp := array[i]
	// child starts as the left child of i
	child := 2*i + 1
	for child < n {
		// if right child is within heap (bounds checking)
		// and the right child is larger than the left
		if child+1 < n && array[child].Compare(array[child+1]) < 0 {
			// right child exists and is bigger
			// so set child to the right child
			child++
		}
		// Is the largest child larger than the node we're trying to move down?
		if array[child].Compare(temp) > 0 {
			array[i] = array[child] // overwrite that node with the child
			i = child               // move the current node to the child
			child = 2*i + 1         // and look at its left child
		} else {
			// if temp isn't bigger than the child
			// then it's in the right place and we can stop
			break
		}
	}
	array[i] = temp
}
