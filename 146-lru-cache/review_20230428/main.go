package main

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	c := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		// add
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		node = initDLinkedNode(key, value)
		// add
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			delete(this.cache, tail.key)
			this.size--
		}
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next, node.next.prev = node.next, node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
