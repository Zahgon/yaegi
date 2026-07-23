package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { _ = "STUB: not implemented"; return 0 }
func (h IntHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (h IntHeap) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *IntHeap) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (h *IntHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	fmt.Println("h:", h)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
