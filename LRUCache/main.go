import (
	"fmt"
)

type LRUCache struct {
	cacheMap map[int]*cacheNode
	capacity int
	head     *cacheNode
	tail     *cacheNode
}

type cacheNode struct {
	key   int
	value int
	prev  *cacheNode
	next  *cacheNode
}

func Constructor(capacity int) LRUCache {
	lruCache := new(LRUCache)
	lruCache.capacity = capacity
	lruCache.cacheMap = make(map[int]*cacheNode)
	return *lruCache
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}

	this.delete(key)
	this.insert(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	_, ok := this.cacheMap[key]
	if len(this.cacheMap) == this.capacity && !ok {
		this.delete(this.head.key)
	}
	this.delete(key)
	node := &cacheNode{
		key:   key,
		value: value,
	}
	this.insert(node)
}

func (this *LRUCache) delete(key int) {
	node, ok := this.cacheMap[key]
	if !ok {
		return
	}
	delete(this.cacheMap, key)
	if this.head == nil {
		return
	}
	if this.head.key == key {
		this.head = this.head.next
		if this.head != nil {
			this.head.prev = nil
		}
		if this.tail != nil && this.tail.key == this.head.key {
			this.tail = nil
		}
		return
	}
	if this.tail == nil {
		return
	}
	if this.tail.key == key {
		this.tail = this.tail.prev
		this.tail.next = nil
		if this.head.key == this.tail.key {
			this.tail = nil
		}
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.next = nil
	node.prev = nil
}

func (this *LRUCache) insert(node *cacheNode) {
	node.next = nil
	node.prev = nil
	this.cacheMap[node.key] = node
	if this.head == nil {
		this.head = node
	} else if this.tail == nil {
		this.tail = node
		this.tail.prev = this.head
		this.head.next = this.tail
	} else {
		this.tail.next = node
		node.prev = this.tail
		this.tail = node
	}
}

/*
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */