// © 2017 Alastair Feille
// Licensed under the MIT License
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Comparer interface {
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
	//fedInput := []MyInt{28, 19, 59, 94, 38, 36, 1, 30, 63, 84, 8, 60, 17, 34, 87, 2, 76, 48, 72, 49}
	numbers := make([]Comparer, 20)

	// fill array with random numbers
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(numbers); i++ {
		numbers[i] = MyInt(rand.Intn(100))
		//numbers[i] = fedInput[i]
	}

	fmt.Println("Original:", numbers)
	buildMaxHeap(numbers)
	fmt.Println("Heap:", numbers)
	heapsort(numbers)
	fmt.Println("Sorted:", numbers)
	fmt.Println()

	// Names
	names := []string{"Brandee", "Nelida", "Jaqueline", "Candyce", "Wayne", "Anissa", "Randal", "Milton", "Manda", "Pasquale", "Alpha", "Destiny", "Romaine", "Waneta", "Claudio", "Arnulfo", "Yukiko", "Barbra", "Judie", "Larry"}

	nodes := make([]Comparer, len(names))
	for i, d := range names {
		nodes[i] = Node{name: d}
	}

	fmt.Println("Original:", nodes)
	buildMaxHeap(nodes)
	fmt.Println("Heap:", nodes)
	heapsort(nodes)
	fmt.Println("Sorted:", nodes)
}

func heapsort(array []Comparer) {
	buildMaxHeap(array)
	heapsize := len(array)
	for i := len(array) - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		heapsize--
		maxHeapify(array, heapsize, 0)
	}
}

func buildMaxHeap(array []Comparer) {
	for i := int(math.Floor(float64(len(array)) / 2)); i >= 0; i-- {
		maxHeapify(array, len(array), i)
	}
}

func maxHeapify(array []Comparer, heapsize int, parent int) {
	for {
		leftChild := (2 * parent) + 1
		rightChild := (2 * parent) + 2

		largest := parent
		if leftChild < heapsize && array[leftChild].Compare(array[largest]) > 0 {
			largest = leftChild
		}

		if rightChild < heapsize && array[rightChild].Compare(array[largest]) > 0 {
			largest = rightChild
		}

		if largest == parent {
			break
		}
		// swap
		array[parent], array[largest] = array[largest], array[parent]
		parent = largest
	}
}
