package tree

//****************************************116. 填充每个节点的下一个右侧节点指针****************************************
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	length := len(queue)
	for length != 0 {
		var preNode *Node
		for i := 0; i < length; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if preNode != nil {
				preNode.Next = node
			}
			preNode = node
			queue = queue[1:]
		}
		length = len(queue)
	}
	return root
}
