package main

import (
	"container/list"
	"fmt"
)
/*
Set operation:
1. If the key already exists, update the key with the new value
2. If the key doesn't exist, the 2 cases can happen
2a. There is enough space in the dictionary, add key, value to dictionary in the set operation
2b. There is not enough space, need to remove the value that is the most ancient from both the
dictionary and linked list that stores this information and add the new value at the head

Get operation:
1. If key doesn't exist, return null
2. If key exists:
2a. Return the value from the key
2b. Travel to the node in the linked list, move to front
 */
type lruCache struct {
	bag map[int]string
	position *list.List
	maxSize int
}

func New(maxSize int) *lruCache {
	return &lruCache{
		bag: make(map[int]string),
		position: list.New(),
		maxSize: maxSize,
	}
}

func (l *lruCache) findElement(key int) *list.Element {
	for e := l.position.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == key {
			return e
		}
	}
	return nil
}

func (l *lruCache) set(key int, value string) {
	_, ok := l.bag[key]
	if ok {
		l.bag[key] = value
		return
	}
	if len(l.bag) >= l.maxSize {
		// remove the last element
		e := l.position.Back()
		lastKey := l.position.Remove(e).(int)
		delete(l.bag, lastKey)
	}
	// add new key to the front
	l.position.PushFront(key)
	l.bag[key] = value
}

func (l *lruCache) get(key int) string {
	val, ok := l.bag[key]
	if !ok {
		return ""
	}
	currElement := l.findElement(key)
	l.position.MoveToFront(currElement)
	return val
}

func main() {
	fmt.Printf("lru_cache_linkedlist \n")
	l := New(3)
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

