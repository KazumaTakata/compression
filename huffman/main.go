package main

import (
	"container/heap"
    "fmt"
)

type Node struct {
    left *Node
    right *Node
    isTerm bool
    value byte
    priority uint
}

type PriorityQueue []*Node

//https://stackoverflow.com/a/24961661/4624070
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
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
	old[n-1] = nil  // avoid memory leak
	*pq = old[0 : n-1]
	return item
}


func main(){





	pq := PriorityQueue{}
    pq = append(pq, &Node{isTerm: true, value: 'e', priority: 3} ) 
    pq = append(pq, &Node{isTerm: true, value: '2', priority: 22} ) 
    pq = append(pq, &Node{isTerm: true, value: 'a', priority: 222} ) 
    pq = append(pq, &Node{isTerm: true, value: '2', priority: 10} ) 





	heap.Init(&pq)


    item := heap.Pop(&pq).(*Node)

    fmt.Printf("%v", item)


    item2 := heap.Pop(&pq).(*Node)

    fmt.Printf("%v", item2)

}
