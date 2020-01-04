package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
)

type Node struct {
	left     *Node
	right    *Node
	isTerm   bool
	value    byte
	priority int
}

type PriorityQueue []*Node

//https://stackoverflow.com/a/24961661/4624070
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Node)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

func main() {

	freq_histo := make(map[byte]int)

	input_data, _ := ioutil.ReadFile("sample.txt")

	for _, d := range input_data {
		if count, ok := freq_histo[d]; ok {
			freq_histo[d] = count + 1
		} else {
			freq_histo[d] = 1
		}
	}

	pq := PriorityQueue{}

	for ch, count := range freq_histo {

		pq = append(pq, &Node{isTerm: true, value: byte(ch), priority: count})

	}

	heap.Init(&pq)

	for len(pq) > 1 {

		item := heap.Pop(&pq).(*Node)
		item2 := heap.Pop(&pq).(*Node)
		priority := item.priority + item2.priority
		item3 := Node{isTerm: false, priority: priority, left: item, right: item2}
		heap.Push(&pq, &item3)
	}

	root := heap.Pop(&pq).(*Node)


    printTree(root)

}

func printTree(node *Node) {

	fmt.Printf("(")
	if node.left != nil {
		printTree(node.left)
	}
	fmt.Printf("[%c,%d]", node.value, node.priority)
	if node.right != nil {
		printTree(node.right)

	}
	fmt.Printf(")")

}
