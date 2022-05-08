package rtda

// Stack 虚拟机栈
type Stack struct {
	maxSize uint
	size    uint
	// 链表方式连接栈帧
	top *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// push 压入栈帧 方法调用
func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if this.top != nil {
		frame.lower = this.top
	}
	this.top = frame
	this.size++
}

// pop 弹出栈 方法返回
func (this *Stack) pop() *Frame {
	if this.top == nil {
		panic("jvm stack is empty!")
	}
	top := this.top
	this.top = top.lower
	top.lower = nil
	this.size--
	return top
}

// peek 查看栈顶
func (this *Stack) peek() *Frame {
	if this.top == nil {
		panic("null stack")
	}
	return this.top
}
