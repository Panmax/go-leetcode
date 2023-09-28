package main

type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode
}

type DLinkNode struct {
	k, v       int
	prev, next *DLinkNode
}

func initDLinkNode(k, v int) *DLinkNode {
	return &DLinkNode{k: k, v: v}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*DLinkNode),
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.v = value
		this.moveToHead(node)
	} else {
		node = initDLinkNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			node = this.removeTail()
			delete(this.cache, node.k)
			this.size--
		}
	}
}

func (this *LRUCache) moveToHead(node *DLinkNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *DLinkNode) {
	node.prev, node.next = this.head, this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next, node.prev = nil, nil
}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
