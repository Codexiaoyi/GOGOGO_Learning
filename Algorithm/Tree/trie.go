package tree

//节点由map和一个是否为结尾的标志量组成

type Trie struct {
	root *node
}

type node struct {
	children map[rune]*node
	isEnd    bool
}

func Constructor() Trie {
	return Trie{root: &node{children: make(map[rune]*node), isEnd: false}}
}

func (this *Trie) Insert(word string) {
	root := this.root
	for _, w := range word {
		if root.children[w] == nil {
			root.children[w] = &node{children: make(map[rune]*node), isEnd: false}
		}
		root = root.children[w]
	}
	//最后一个节点标记为结尾
	root.isEnd = true
}

func (this *Trie) Search(word string) bool {
	//每个字母都能搜索到，并且最后一个字母是在结尾，就是搜索成功
	cur := this.root
	for _, w := range word {
		if cur.children[w] == nil {
			return false
		}
		cur = cur.children[w]
	}
	if cur.isEnd {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, w := range prefix {
		if cur.children[w] == nil {
			return false
		}
		cur = cur.children[w]
	}
	return true
}
