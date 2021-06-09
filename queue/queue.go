package queue

import "github.com/fishwin/mylib/single_list"

type Queue struct {
	list *single_list.SingleList
}

func NewQueue() *Queue {
	q := &Queue{
		list: single_list.NewSingleList(),
	}

	return q
}

func (q *Queue) Size() uint {
	return q.list.Length()
}

func (q *Queue) Enqueue(data interface{}) bool {
	return q.list.Append(&single_list.SingleNode{Data: data})
}

func (q *Queue) Dequeue() interface{} {
	node := q.list.Get(0)
	if node == nil {
		return nil
	}
	q.list.Delete(0)
	return node.Data
}

func (q *Queue) Peek() interface{} {
	node := q.list.Get(0)
	if node == nil {
		return nil
	}
	return node.Data
}
