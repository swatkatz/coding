package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

/**
What does ring have?
Keeps track of the list of nodes
Keeps track of the number of replicas that we want
values are strings
 */

type hashFn func(data []byte) uint32

type ring struct {
	replicas int
	hashFunction hashFn
	sortedKeys []int
	hashMap map[int]string
}

// can also take in a hash function, but kiss for now
func newRing(replicas int) *ring {
	return &ring{
		replicas:     replicas,
		hashFunction: crc32.ChecksumIEEE,
		sortedKeys:   make([]int, 0),
		hashMap:      make(map[int]string),
	}
}

func (r *ring) isEmpty() bool {
	return len(r.sortedKeys) == 0
}

func (r *ring) add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < r.replicas; i++ {
			hash := int(r.hashFunction([]byte(strconv.Itoa(i) + key)))
			r.hashMap[hash] = key
			r.sortedKeys = append(r.sortedKeys, hash)
		}
	}
	sort.Ints(r.sortedKeys)
}

func (r *ring) get(key string) string {
	if r.isEmpty() {
		return ""
	}
	hash := int(r.hashFunction([]byte(key)))

	// binary search
	idx := sort.Search(
		len(r.sortedKeys),
		func(i int) bool { return r.sortedKeys[i] >= hash },
	)

	// circled back
	if idx == len(r.sortedKeys) {
		idx = 0
	}
	return r.hashMap[r.sortedKeys[idx]]
}

func (r *ring) remove(key string) {
	if r.isEmpty() {
		return
	}
	hash := int(r.hashFunction([]byte(key)))
	// binary search
	idx := sort.Search(
		len(r.sortedKeys),
		func(i int) bool { return r.sortedKeys[i] >= hash },
	)
	// couldn't find the key to remove
	if idx == len(r.sortedKeys) {
		return
	}
	delete(r.hashMap, r.sortedKeys[idx])
	r.sortedKeys = append(r.sortedKeys[:idx], r.sortedKeys[idx + 1:]...)
}

func main() {
	r := newRing(1)
	r.add("A")
	r.add("B")
	r.add("C")
	fmt.Printf("BEFORE: %v", r.get("FOO") + "\n")
	fmt.Printf("BEFORE: %v", r.get("BAR") + "\n")
	r.remove("C")
	fmt.Printf(r.get("BAR") + "\n")
	fmt.Printf(r.get("FOO") + "\n")
}
