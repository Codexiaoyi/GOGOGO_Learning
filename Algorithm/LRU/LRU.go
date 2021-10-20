package lru

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
func LRU(operators [][]int, k int) []int {
	// write code here
	lru := newLru(k)
	res := make([]int, 0)
	for i := 0; i < len(operators); i++ {
		if operators[i][0] == 1 {
			lru.set(operators[i][1], operators[i][2])
		}
		if operators[i][0] == 2 {
			res = append(res, lru.get(operators[i][1]))
		}
	}
	return res
}

type Lru struct {
	max     int
	current int
	root    *Node
	mcache  map[int]*Node
}

type Node struct {
	key, value int
	pre, next  *Node
}

func newLru(max int) *Lru {
	lru := &Lru{
		max:     max,
		current: 0,
		root:    &Node{},
		mcache:  make(map[int]*Node),
	}
	lru.root.next = lru.root
	lru.root.pre = lru.root
	return lru
}

func (l *Lru) set(key, value int) {
	if n, ok := l.mcache[key]; ok {
		n.value = value
		return
	}
	node := &Node{key: key, value: value}
	l.mcache[key] = node
	l.pushToHead(node)
	l.current++
	if l.current > l.max {
		tail := l.removeAtTail()
		delete(l.mcache, tail.key)
	}
}

func (l *Lru) get(key int) int {
	if node, ok := l.mcache[key]; ok {
		l.moveToHead(node)
		return node.value
	}
	return -1
}

func (l *Lru) pushToHead(node *Node) {
	l.root.next.pre = node
	node.next = l.root.next
	node.pre = l.root
	l.root.next = node
}

func (l *Lru) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
	node.next = nil
	node.pre = nil
}

func (l *Lru) removeAtTail() *Node {
	tail := l.root.pre
	l.remove(tail)
	return tail
}

func (l *Lru) moveToHead(node *Node) {
	l.remove(node)
	l.pushToHead(node)
}
