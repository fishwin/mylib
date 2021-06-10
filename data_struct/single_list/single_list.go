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
	head  *SingleNode
	tail  *SingleNode
	size  uint
}

func NewSingleList() *SingleList {
	sl := &SingleList{
		mutex: &sync.RWMutex{},
		head:  nil,
		tail:  nil,
		size:  0,
	}

	return sl
}

func (list *SingleList) Append(node *SingleNode) bool {
	if node == nil {
		return false
	}
	list.mutex.Lock()
	defer list.mutex.Unlock()
	if list.size == 0 {
		list.head = node
		list.tail = node
		list.size = 1
		return true
	}

	tail := list.tail
	tail.Next = node
	list.tail = node
	list.size += 1
	return true
}

func (list *SingleList) Insert(index uint, node *SingleNode) bool {
	if node == nil {
		return false
	}

	if index > list.size {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		node.Next = list.head
		list.head = node
		list.size += 1
		return true
	}
	var i uint
	ptr := list.head
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next
	ptr.Next = node
	node.Next = next
	list.size += 1
	return true
}

func (list *SingleList) Delete(index uint) bool {
	if list == nil || list.size == 0 || index > list.size-1 {
		return false
	}

	list.mutex.Lock()
	defer list.mutex.Unlock()

	if index == 0 {
		head := list.head.Next
		list.head = head
		if list.size == 1 {
			list.tail = nil
		}
		list.size -= 1
		return true
	}

	ptr := list.head
	var i uint
	for i = 1; i < index; i++ {
		ptr = ptr.Next
	}
	next := ptr.Next

	ptr.Next = next.Next
	if index == list.size-1 {
		list.tail = ptr
	}
	list.size -= 1
	return true
}

func (list *SingleList) Get(index uint) *SingleNode {
	if list == nil || list.size == 0 || index > list.size-1 {
		return nil
	}

	list.mutex.RLock()
	defer list.mutex.RUnlock()

	if index == 0 {
		return list.head
	}
	node := list.head
	var i uint
	for i = 0; i < index; i++ {
		node = node.Next
	}
	return node
}

func (list *SingleList) Length() uint {
	return list.size
}

func (list *SingleList) Display() {
	if list == nil || list.size == 0 {
		return
	}
	list.mutex.RLock()
	defer list.mutex.RUnlock()
	ptr := list.head
	var i uint
	for i = 0; i < list.size; i++ {
		fmt.Printf("data is %v\n", ptr.Data)
		ptr = ptr.Next
	}
}
