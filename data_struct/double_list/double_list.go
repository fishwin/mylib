package double_list

import (
	"fmt"
	"sync"
)

type DoubleNode struct {
	Data interface{}
	Prev *DoubleNode
	Next *DoubleNode
}

type DoubleList struct {
	mutex *sync.RWMutex
	Size  uint
	Head  *DoubleNode
	Tail  *DoubleNode
}

func NewDoubleList() *DoubleList {
	dl := &DoubleList{
		mutex: &sync.RWMutex{},
		Size:  0,
		Head:  nil,
		Tail:  nil,
	}

	return dl
}

func (list *DoubleList) Get(index uint) *DoubleNode {
	if list.Size == 0 || index > list.Size-1 {
		return nil
	}
	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 1; i <= index; i++ {
		node = node.Next
	}
	return node
}

func (list *DoubleList) Append(node *DoubleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		node.Next = nil
		node.Prev = nil
	} else {
		node.Prev = list.Tail
		node.Next = nil
		list.Tail.Next = node
		list.Tail = node
	}
	list.Size++
	return true
}

func (list *DoubleList) Insert(index uint, node *DoubleNode) bool {
	if index > list.Size || node == nil {
		return false
	}

	if index == list.Size {
		return list.Append(node)
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Head.Prev = nil
		list.Size++
		return true
	}

	nextNode := list.Get(index)
	node.Prev = nextNode.Prev
	node.Next = nextNode
	nextNode.Prev.Next = node
	nextNode.Prev = node
	list.Size++
	return true
}

func (list *DoubleList) Delete(index uint) bool {
	if index+1 > list.Size {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()
	if index == 0 {
		if list.Size == 1 {
			list.Head = nil
			list.Tail = nil
		} else {
			list.Head.Next.Prev = nil
			list.Head = list.Head.Next
		}
		list.Size--
		return true
	}
	if index+1 == list.Size {
		list.Tail.Prev.Next = nil
		list.Tail = list.Tail.Prev
		list.Size--
		return true
	}

	node := list.Get(index)
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	list.Size--
	return true
}

func (list *DoubleList) Display() {
	if list == nil || list.Size == 0 {
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	ptr := list.Head
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}

func (list *DoubleList) Reverse() {
	if list == nil || list.Size == 0 {
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	ptr := list.Tail
	for ptr != nil {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Prev
	}
}
