package deque

type Node struct {
	v          interface{}
	prev, next *Node
}

type Deque struct {
	head, tail *Node
	size       int
}

func CreateDeque() Deque {
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return Deque{head, tail, 0}
}

func (d *Deque) PushHead(v interface{}) {
	t := &Node{v: v}
	t.prev = d.head
	t.next = d.head.next
	d.head.next = t
	t.next.prev = t
	d.size++
}

func (d *Deque) PopHead() interface{} {
	t := d.head.next
	d.head.next = t.next
	t.next.prev = d.head
	t.prev, t.next = nil, nil
	d.size--
	return t.v
}

func (d *Deque) PushTail(v interface{}) {
	t := &Node{v: v}
	t.prev = d.tail.prev
	t.next = d.tail
	t.prev.next = t
	d.tail.prev = t
	d.size++
}

func (d *Deque) PopTail() interface{} {
	t := d.tail.prev
	d.tail.prev = t.prev
	t.prev.next = d.tail
	t.prev, t.next = nil, nil
	d.size--
	return t.v
}

func (d *Deque) IsEmpty() bool {
	return d.size == 0
}
