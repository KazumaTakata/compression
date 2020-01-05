package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"unicode"
)

type Node struct {
	left     *Node
	right    *Node
	isTerm   bool
	value    uint
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

		pq = append(pq, &Node{isTerm: true, value: uint(ch), priority: count})

	}

	var eof uint = 256

	pq = append(pq, &Node{isTerm: true, value: eof, priority: 1})

	heap.Init(&pq)

	for len(pq) > 1 {

		item := heap.Pop(&pq).(*Node)
		item2 := heap.Pop(&pq).(*Node)
		priority := item.priority + item2.priority
		item3 := Node{isTerm: false, priority: priority, left: item, right: item2}
		heap.Push(&pq, &item3)
	}

	root := heap.Pop(&pq).(*Node)

	encode_map := make(map[uint]string)

	walk_tree(root, encode_map, "")

	for key, value := range encode_map {
		if unicode.IsSpace(rune(key)) {
			switch key {
			case ' ':
				fmt.Printf("%s:%s\n", "<space>", value)
			case '\n':
				fmt.Printf("%s:%s\n", "<newline>", value)

			}
		} else {
			fmt.Printf("%c:%s\n", key, value)
		}
	}

	encoded_data := ""

	for _, data := range input_data {
		encoded := encode_map[uint(data)]
		encoded_data += encoded
	}
	encoded_data += encode_map[eof]

	fmt.Printf("%s\n", encoded_data)


    var decoded uint
    for len(encoded_data) > 0 {
         encoded_data, decoded =  decode(encoded_data, root)
            
         if decoded == eof {
            break
         }
         fmt.Printf("%c", decoded)
    }
}

func decode(encoded_data string, node *Node) (string, uint) {
	if node.isTerm {
		return encoded_data, node.value
	} else {

		bit := encoded_data[0]

		if bit == '1' {
			return decode(encoded_data[1:], node.left)
		}

		if bit == '0' {
			return decode(encoded_data[1:], node.right)
		}
	}

    return "", 0 
}

func walk_tree(node *Node, encode_map map[uint]string, prefix string) {

	if node.left != nil {
		walk_tree(node.left, encode_map, prefix+"1")
	}
	if node.right != nil {
		walk_tree(node.right, encode_map, prefix+"0")
	}

	if node.isTerm {
		encode_map[uint(node.value)] = prefix
	}

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
