package test

import (
	"container/heap"
	"fmt"
	"testing"
)

type stu struct {
	name string
	age  int
}

type Stu []stu

func (s *Stu) Len() int {
	return len(*s)
}

func (s *Stu) Less(i, j int) bool {
	return (*s)[i].age < (*s)[j].age
}

func (s *Stu) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Stu) Push(x interface{}) {
	*s = append(*s, x.(stu))
}

func (s *Stu) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[:n-1]
	return x
}

func TestPq(t *testing.T) {
	student := &Stu{{"Amy", 21}, {"Dav", 15}, {"Spo", 22}, {"Reb", 11}}
	heap.Init(student)
	one := stu{"hund", 9}
	heap.Push(student, one)
	for student.Len() > 0 {
		fmt.Printf("%v\n", heap.Pop(student))
	}
}
