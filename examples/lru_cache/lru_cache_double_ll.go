package main

import "fmt"

type dll struct {
	head *node
	tail *node
}

func (d *dll) moveToFront(n *node) {
	d.removeNode(n)
	d.addNode(n)
}

func (d *dll) popTail() *node {
	e := d.tail.prev
	d.removeNode(e)
	return e
}

// add node to head
func (d *dll) addNode(n *node) {
	e := d.head.next
	d.head.next = n
	n.prev = d.head

	n.next = e
	e.prev = n
}

func (d *dll) removeNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

type node struct {
	key int
	val string
	prev *node
	next *node
}

func NewDll() *dll {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return &dll{
		head: head,
		tail: tail,
	}
}

type lru struct {
	bag map[int]*node
	ll *dll
	maxLength int
}

func (l *lru) set(key int, value string) {
	n, ok := l.bag[key]
	if ok {
		n.val = value
		l.bag[key] = n
		return
	}

	if len(l.bag) >= l.maxLength {
		lastEle := l.ll.popTail()
		delete(l.bag, lastEle.key)
	}
	// add new node to the front
	newNode := &node{key: key, val: value}
	l.ll.addNode(newNode)
	l.bag[key] = newNode
}


func (l *lru) get(key int) string {
	n, ok := l.bag[key]
	if !ok {
		return ""
	}
	l.ll.moveToFront(n)
	return n.val
}

func NewLru(maxLength int) *lru {
	return &lru{
		bag: make(map[int]*node),
		ll: NewDll(),
		maxLength: maxLength,
	}
}

func main() {
	fmt.Printf("lru_cache_double \n")
	l := NewLru(3)
	l.set(10, "foo")
	l.set(20, "bar")
	l.set(21, "baz")
	l.get(10)
	fmt.Printf("l.get(10): %v \n", l.get(10))
	l.set(22, "abc")
	fmt.Printf("l.get(20): %v \n", l.get(20))
	fmt.Printf("l.get(21): %v \n", l.get(21))
	fmt.Printf("l.get(22): %v \n", l.get(22))
}
