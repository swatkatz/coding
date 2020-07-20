package main

import "fmt"

type node struct {
	data int
	next *node
}

type list struct {
	root *node
	count int
}

func (l *list) insert(data int) {
	var n *node
	if l.root == nil {
		n = &node{
			data: data,
		}
	} else {
		n = &node{
			data: data,
			next: l.root,
		}
	}
	l.root = n
	l.count += 1
}

func (l *list) print() {
	printUtil(l.root)
	fmt.Printf("\n")
}

func (l *list) len() int {
	return l.count
}

func (l *list) search(data int) *node {
	return searchUtil(l.root, data)
}

func (l *list) delete(data int) {
	// Get the current node that has the data
	curr := searchUtil(l.root, data)
	// node with data doesn't exist
	if curr == nil {
		return
	}
	prev := searchPrevUtil(l.root, data)
	l.count -= 1
	// root
	if prev == nil {
		l.root = curr.next
		return
	}
	prev.next = curr.next
}

func searchUtil(n *node, data int) *node {
	if n == nil {
		return nil
	}
	if n.data == data {
		return n
	}
	return searchUtil(n.next, data)
}

func searchPrevUtil(n *node, data int) *node {
	if n == nil || n.next == nil {
		return nil
	}
	if n.next.data == data {
		return n
	}
	return searchPrevUtil(n.next, data)
}

func printUtil(n *node) {
	if n != nil {
		fmt.Printf("%v", n.data)
		if n.next != nil {
			fmt.Printf(" -> ")
			printUtil(n.next)
		}
	}
}

func main() {
	l := &list{
		root:  nil,
		count: 0,
	}
	l.insert(1)
	l.insert(2)
	l.insert(3)

	l.print()
	fmt.Printf("len: %v \n", l.len())

	l.delete(2)
	l.print()
	fmt.Printf("len: %v \n", l.len())

	l.delete(3)
	l.print()
	fmt.Printf("len: %v \n", l.len())

	l.delete(1)
	l.print()
	fmt.Printf("len: %v \n", l.len())
}
