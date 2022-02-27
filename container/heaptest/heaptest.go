package heaptest

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func IntHeapTest() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h).(int))
	}
}

//self struct heap test

type People struct {
	name  string
	age   int
	grade int
}

type PeopleHeap []People

func (p PeopleHeap) Len() int           { return len(p) }
func (p PeopleHeap) Less(i, j int) bool { return p[i].age < p[j].age }
func (p PeopleHeap) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func (p *PeopleHeap) Push(x interface{}) {
	*p = append(*p, x.(People))
}

func (p *PeopleHeap) Pop() interface{} {
	old := *p
	n := len(*p)
	x := old[n-1]
	*p = old[0 : n-1]
	return x
}

func PrintPeople(x People) {
	fmt.Printf("%s, %d, %d\n", x.name, x.age, x.grade)
}

func PeopleHeapTest() {
	p := &PeopleHeap{People{
		name:  "Milittle",
		age:   10,
		grade: 100,
	},
		People{
			name:  "Fire",
			age:   12,
			grade: 99,
		}}
	heap.Init(p)
	heap.Push(p, People{
		name:  "Little",
		age:   5,
		grade: 120,
	})

	PrintPeople((*p)[0])

	for p.Len() > 0 {
		PrintPeople(heap.Pop(p).(People))
	}
}
