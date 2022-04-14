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

//****************************************106. 从中序与后序遍历序列构造二叉树***************************************
func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderLen := len(inorder)
	if inorderLen == 1 {
		return &TreeNode{Val: inorder[0]}
	}
	inorderIndex := 0
	//后序最后一个是根节点
	for i := 0; i < inorderLen; i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			inorderIndex = i
			break
		}
	}
	root := &TreeNode{Val: inorder[inorderIndex]}
	if inorderIndex > 0 {
		leftInorder := inorder[:inorderIndex]
		root.Left = buildTree(leftInorder, postorder[:inorderIndex])
	}
	if inorderIndex < inorderLen-1 {
		rightInorder := inorder[inorderIndex+1:]
		root.Right = buildTree(rightInorder, postorder[inorderIndex:len(postorder)-1])
	}
	return root
}

//****************************************589. N 叉树的前序遍历***************************************
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	indexMap := make(map[*Node]int)
	stack := make([]*Node, 0)
	stack = append(stack, root)
	for root != nil || len(stack) > 0 {
		//遍历所有的左侧根节点
		for root != nil {
			res = append(res, root.Val)
			if len(root.Children) == 0 {
				//叶子节点，不入栈不迭代
				break
			}
			indexMap[root] = 0
			stack = append(stack, root)
			root = root.Children[0]
		}

		node := stack[len(stack)-1]
		if indexMap[node] < len(node.Children)-1 {
			indexMap[node]++
			root = node.Children[indexMap[node]]
		} else {
			stack = stack[:len(stack)-1]
			root = nil
		}
	}
	return res
}
