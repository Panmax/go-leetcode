package main

type LFUCache struct {
	cache    map[int]*node
	freqMap  map[int]*list
	minFreq  int
	capacity int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:    make(map[int]*node),
		freqMap:  make(map[int]*list),
		capacity: capacity,
	}
}

func (this *LFUCache) Get(key int) int {
	if n, ok := this.cache[key]; ok {
		this.incrFreq(n)
		return n.value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if n, ok := this.cache[key]; ok {
		n.value = value
		this.incrFreq(n)
		return
	}
	n := newNode(key, value)
	if len(this.cache) == this.capacity {
		l := this.freqMap[this.minFreq]
		removeNode := l.removeTail()
		delete(this.cache, removeNode.key)
	}
	n.freq = 1
	this.cache[key] = n
	this.minFreq = 1
	this.addNode(n)
}

func (this *LFUCache) incrFreq(n *node) {
	l := this.freqMap[n.freq]
	l.removeNode(n)
	if l.empty() && n.freq == this.minFreq {
		this.minFreq++
	}
	n.freq++
	this.addNode(n)
}

func (this *LFUCache) addNode(n *node) {
	l := this.freqMap[n.freq]
	if l == nil {
		l = newList()
		this.freqMap[n.freq] = l
	}
	l.addFirst(n)
}

type node struct {
	key, value int
	prev, next *node
	freq       int
}

func newNode(k, v int) *node {
	return &node{
		key:   k,
		value: v,
	}
}

type list struct {
	head, tail *node
}

func (l *list) removeNode(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *list) empty() bool {
	return l.head.next == l.tail
}

func (l *list) addFirst(n *node) {
	n.prev, n.next = l.head, l.head.next
	l.head.next.prev = n
	l.head.next = n
}

func (l *list) removeTail() *node {
	n := l.tail.prev
	l.removeNode(n)
	return n
}

func newList() *list {
	l := &list{
		head: newNode(0, 0),
		tail: newNode(0, 0),
	}
	l.head.next = l.tail
	l.head.prev = l.head
	return l
}
