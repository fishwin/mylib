package main

import (
	"fmt"

	"github.com/fishwin/mylib/data_struct/double_list"
)

func main() {
	dl := double_list.NewDoubleList()
	dl.Append(&double_list.DoubleNode{Data: 1})
	dl.Append(&double_list.DoubleNode{Data: 2})
	dl.Append(&double_list.DoubleNode{Data: 3})
	dl.Append(&double_list.DoubleNode{Data: 4})
	dl.Append(&double_list.DoubleNode{Data: 5})

	dl.Display()
	dl.Reverse()

	dl.Delete(2)

	dl.Display()
	dl.Reverse()

	dl.Insert(4, &double_list.DoubleNode{Data: 111})

	dl.Display()
	dl.Reverse()

	fmt.Println(dl.Get(3))
}
