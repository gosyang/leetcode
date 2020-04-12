// Package graph - package brief introduce
/* Copyright 2019 All Rights Reserved. */
/* xx - file brief introduce */
/*
modification history
-----------------------------
2020/4/12 8:57 PM, by gosyang, create
*/
/*
DESCRIPTION
xx
*/
package design

type LRUCache struct {
	// capacity of LRUCache
	capacity int
	// main cache of LRUCache
	cache map[int]*LRUListNode

	// help LRUListNode, real LRUListNode are between head and tail
	head *LRUListNode
	tail *LRUListNode
}

type LRUListNode struct {
	key int
	value int

	pre *LRUListNode
	next *LRUListNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		cache: make(map[int]*LRUListNode, capacity),
		head: new(LRUListNode),
		tail: new(LRUListNode),
	}

	lru.head.next = lru.tail
	lru.tail.next = lru.head

	return lru
}


// setHeader insert new node between head and old first real node
// like: this.head <-> node <-> oldFirstNode
func (this *LRUCache) setHeader(node *LRUListNode) {
	node.next = this.head.next
	node.pre = this.head

	node.next.pre = node
	this.head.next = node
}

// deleteNode delete node from LRUCache NodeList
func (this *LRUCache) deleteNode(node *LRUListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.deleteNode(node)
	this.setHeader(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		if len(this.cache) >= this.capacity {
			toDelete := this.tail.pre
			this.deleteNode(toDelete)
			delete(this.cache, toDelete.key)
		}
	} else {
		this.deleteNode(node)
		delete(this.cache, node.key)
	}

	newNode := &LRUListNode{
		key: key,
		value: value,
	}
	this.setHeader(newNode)
	this.cache[key] = newNode
}
