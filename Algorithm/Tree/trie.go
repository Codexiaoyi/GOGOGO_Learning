package tree

type Trie struct {
	root *node
}

type node struct {
	isEnd  bool
	childs map[byte]*node
}

func Constructor() Trie {
	return Trie{root: &node{childs: make(map[byte]*node)}}
}

func (this *Trie) Insert(word string) {
	cur := this.root
	for _, b := range word {
		if node, ok := cur.childs[byte(b)]; ok {
			cur = node
			continue
		}
		newNode := &node{childs: make(map[byte]*node)}
		cur.childs[byte(b)] = newNode
		cur = newNode
	}
	cur.isEnd = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, b := range word {
		if node, ok := cur.childs[byte(b)]; ok {
			cur = node
			continue
		}
		return false
	}
	if cur.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, b := range prefix {
		if node, ok := cur.childs[byte(b)]; ok {
			cur = node
			continue
		}
		return false
	}
	return true
}
