package offer

//****************************************用两个栈实现队列****************************************

type CQueue struct {
	stack_set    []int
	stack_delete []int
}

func Constructor() CQueue {
	return CQueue{stack_set: make([]int, 0), stack_delete: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	this.stack_set = append(this.stack_set, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.stack_delete) == 0 {
		sl := len(this.stack_set)
		if sl == 0 {
			return -1
		}

		for i := 0; i < sl; i++ {
			if len(this.stack_set) == 0 {
				this.stack_delete = append(this.stack_delete, -1)
			} else {
				this.stack_delete = append(this.stack_delete, this.stack_set[len(this.stack_set)-1])
				this.stack_set = this.stack_set[:len(this.stack_set)-1]
			}
		}
	}

	res := this.stack_delete[len(this.stack_delete)-1]
	this.stack_delete = this.stack_delete[:len(this.stack_delete)-1]
	return res
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
