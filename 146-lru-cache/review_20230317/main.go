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
		head:     initDLinkedNode(0, 0),
		tail:     initDLinkedNode(0, 0),
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	return c
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		this.moveToHead(node)
		return node.value
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		node = initDLinkedNode(key, value)
		// add
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			tail := this.removeTail()
			// delete
			delete(this.cache, tail.key)
			this.size--
		}
	} else {
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.next = this.head.next
	node.prev = this.head
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
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
