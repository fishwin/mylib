package single_list

import (
	"fmt"
	"sync"
)

type SingleNode struct {
	Data interface{}
	Next *SingleNode
}

type SingleList struct {
	mutex *sync.RWMutex
	Head  *SingleNode
	Tail  *SingleNode
	Size  uint
}

func NewSingleList() *SingleList {
	sl := &SingleList{
		mutex: &sync.RWMutex{},
		Head:  nil,
		Tail:  nil,
		Size:  0,
	}

	return sl
}

func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.Size == 0 {
		list.Head = node
		list.Tail = node
		list.Size = 1
		return true
	}

	tail := list.Tail
	tail.Next = node
	list.Tail = node
	list.Size += 1
	return true
}

func (list *SingleList) Insert(index uint, node *SingleNode) bool {
	if node == nil {
		return false
	}

	if index > list.Size {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		node.Next = list.Head
		list.Head = node
		list.Size += 1
		return true
	}
	var i uint
	ptr := list.Head
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	ptr.Next = node
	node.Next = next
	list.Size += 1
	return true
}

func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		head := list.Head.Next
		list.Head = head
		if list.Size == 1 {
			list.Tail = nil
		}
		list.Size -= 1
		return true
	}

	ptr := list.Head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next

	ptr.Next = next.Next
	if index == list.Size-1 {
		list.Tail = ptr
	}
	list.Size -= 1
	return true
}

func (list *SingleList) Get(index uint) *SingleNode {
	if list == nil || list.Size == 0 || index > list.Size-1 {
		return nil
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()

	if index == 0 {
		return list.Head
	}
	node := list.Head
	var i uint
	for i = 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (list *SingleList) Length() uint {
	return list.Size
}

func (list *SingleList) Display() {
	if list == nil || list.Size == 0 {
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	ptr := list.Head
	var i uint
	for i = 0; i < list.Size; i++ {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}
