//Package stack implements stack of int using linkedlist
package stack

type stackNode struct {
	val  int
	next *stackNode
}

type stack struct {
	head *stackNode
}

func (s *stack) push(val int) {
	newNode := &stackNode{
		val: val,
	}
	if s.head == nil {
		s.head = newNode
	} else {
		newNode.next = s.head
		s.head = newNode
	}
}

func (s *stack) pop() int {
	if s.head == nil {
		return -1
	}
	elem := s.head.val
	s.head = s.head.next
	return elem
}

func main() {
}
