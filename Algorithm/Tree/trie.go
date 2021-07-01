package tree

type Trie struct {
	root *node
}

type node struct {
	isEnd bool
	child map[rune]*node
}

func Constructor() Trie {
	trie := Trie{root: &node{
		isEnd: false,
		child: make(map[rune]*node),
	}}
	return trie
}

func (trie *Trie) Insert(word string) {
	cur := trie.root
	for _, w := range word {
		if cur.child[w] == nil {
			cur.child[w] = &node{
				isEnd: false,
				child: make(map[rune]*node),
			}
		}
		cur = cur.child[w]
	}
	cur.isEnd = true
}

func (trie *Trie) Search(word string) bool {
	cur := trie.root
	for _, w := range word {
		if cur.child[w] == nil {
			return false
		}

		cur = cur.child[w]
	}
	if cur.isEnd {
		return true
	}
	return false
}

func (trie *Trie) StartsWith(prefix string) bool {
	cur := trie.root
	for _, w := range prefix {
		if cur.child[w] == nil {
			return false
		}

		cur = cur.child[w]
	}
	return true
}
