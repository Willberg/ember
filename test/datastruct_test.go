package test

import (
	"fmt"
	dll "github.com/emirpasic/gods/lists/doublylinkedlist"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	"github.com/emirpasic/gods/queues/priorityqueue"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"github.com/emirpasic/gods/utils"
	"testing"
)

func TestLinkedListStack(t *testing.T) {
	stack := lls.New()  // empty
	stack.Push(1)       // 1
	stack.Push(2)       // 1, 2
	stack.Values()      // 2, 1 (LIFO order)
	_, _ = stack.Peek() // 2,true
	_, _ = stack.Pop()  // 2, true
	_, _ = stack.Pop()  // 1, true
	_, _ = stack.Pop()  // nil, false (nothing to pop)
	stack.Push(1)       // 1
	stack.Clear()       // empty
	stack.Empty()       // true
	stack.Size()        // 0
}

func TestQueue(t *testing.T) {
	queue := llq.New()     // empty
	queue.Enqueue(1)       // 1
	queue.Enqueue(2)       // 1, 2
	_ = queue.Values()     // 1, 2 (FIFO order)
	_, _ = queue.Peek()    // 1,true
	_, _ = queue.Dequeue() // 1, true
	_, _ = queue.Dequeue() // 2, true
	_, _ = queue.Dequeue() // nil, false (nothing to deque)
	queue.Enqueue(1)       // 1
	queue.Clear()          // empty
	queue.Empty()          // true
	_ = queue.Size()       // 0
}

func TestDeque(t *testing.T) {
	list := dll.New()
	list.Add("a")                         // ["a"]
	list.Add("c", "b")                    // ["a","c","b"]
	list.Sort(utils.StringComparator)     // ["a","b","c"]
	_, _ = list.Get(0)                    // "a",true
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
	list.Insert(0, "b")                   // ["b"]
	list.Insert(0, "a")                   // ["a","b"]
}

type MyCalendarThree struct {
	cs  [][]int
	cnt int
}

func (this *MyCalendarThree) Book() {
	cs := this.cs
	cs[0][0], cs[0][1] = cs[0][1], cs[0][0]
}

type trie struct {
	children [26]*trie
	isEnd    bool
}

func TestMy(t *testing.T) {
	fmt.Println(findKthLargest([]int{1, 2, 3}, 2))
}

func findKthLargest(nums []int, k int) int {
	pq := priorityqueue.NewWith(utils.IntComparator)
	for _, v := range nums {
		pq.Enqueue(v)
	}
	for k > 1 {
		pq.Dequeue()
		k--
	}
	ans, _ := pq.Dequeue()
	return ans.(int)
}
