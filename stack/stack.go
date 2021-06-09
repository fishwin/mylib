package stack

import "github.com/fishwin/mylib/single_list"

type Stack struct {
	list *single_list.SingleList
}

func NewStack() *Stack {
	return &Stack{list: single_list.NewSingleList()}
}

func (s *Stack) Push(data interface{}) bool {
	node := &single_list.SingleNode{
		Data: data,
	}
	return s.list.Insert(0, node)
}

func (s *Stack) Pop() interface{} {
	node := s.list.Get(0)
	if node != nil {
		s.list.Delete(0)
		return node.Data
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	node := s.list.Get(0)
	if node != nil {
		return node.Data
	}
	return nil
}

func (s *Stack) Length() uint {
	return s.list.Length()
}
