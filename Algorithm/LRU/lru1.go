package lru

type LRUCache struct {
	max    int
	length int
	root   *lruNode
	m      map[int]*lruNode
}

type lruNode struct {
	key   int
	value int
	pre   *lruNode
	next  *lruNode
}

func Constructor(capacity int) LRUCache {
	node := &lruNode{}
	node.pre = node
	node.next = node
	return LRUCache{
		max:  capacity,
		root: node,
		m:    make(map[int]*lruNode),
	}
}

func (this *LRUCache) Get(key int) int {
	//先判断是不是已经有了
	node, ok := this.m[key]
	if !ok {
		//没有就返回-1
		return -1
	}
	//被访问了就移动到队头
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	//先判断是不是已经有了
	if node, ok := this.m[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	newNode := &lruNode{
		key:   key,
		value: value,
	}
	newNode.next = this.root.next
	newNode.pre = this.root
	this.root.next.pre = newNode
	this.root.next = newNode
	this.moveToHead(newNode)

	this.m[key] = newNode
	this.length++

	if this.length > this.max {
		//淘汰
		this.removeWithTail()
	}
}

func (this *LRUCache) moveToHead(node *lruNode) {
	if node == this.root || node.pre == this.root {
		return
	}
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next = this.root.next
	node.pre = this.root
	this.root.next.pre = node
	this.root.next = node
}

func (this *LRUCache) removeWithTail() {
	tail := this.root.pre
	if tail == this.root {
		//空
		return
	}
	tail.pre.next = this.root
	this.root.pre = tail.pre
	tail.next = nil
	tail.pre = nil
	delete(this.m, tail.key)
}
