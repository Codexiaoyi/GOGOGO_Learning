package lru

type LRUCache struct {
	capacity      int
	currentLength int
	root          *node
	m             map[int]*node
}

type node struct {
	key, value int
	pre, next  *node
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity:      capacity,
		currentLength: 0,
		root:          &node{},
		m:             make(map[int]*node),
	}
	//root.next是队列头，root.pre是队列尾部
	lru.root.pre = lru.root
	lru.root.next = lru.root
	return lru
}

func (this *LRUCache) Get(key int) int {
	n, ok := this.m[key]
	if !ok {
		return -1
	}
	if this.root.next != n {
		this.moveToHead(n)
	}
	return n.value
}

func (this *LRUCache) Put(key int, value int) {
	if n, ok := this.m[key]; ok {
		n.value = value
		this.moveToHead(n)
		return
	}
	newNode := &node{key: key, value: value}
	this.moveToHead(newNode)
	this.m[key] = newNode
	this.currentLength++
	if this.currentLength > this.capacity {
		this.removeAtTail()
	}
}

func (this *LRUCache) moveToHead(n *node) {
	if n == this.root.next {
		return
	}
	if n.pre != nil && n.next != nil {
		n.pre.next = n.next
		n.next.pre = n.pre
	}
	temp := this.root.next
	this.root.next = n
	n.pre = this.root
	n.next = temp
	temp.pre = n
}

func (this *LRUCache) removeAtTail() {
	if this.root.pre == this.root {
		return
	}
	tail := this.root.pre
	tail.pre.next = this.root
	this.root.pre = tail.pre
	tail.next = nil
	tail.pre = nil
	delete(this.m, tail.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
