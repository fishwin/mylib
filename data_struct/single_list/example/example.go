package main

import (
	"fmt"

	"github.com/fishwin/mylib/data_struct/single_list"
)

func main() {
	sl := single_list.NewSingleList()
	sl.Append(&single_list.SingleNode{Data: 1})
	sl.Append(&single_list.SingleNode{Data: 2})
	sl.Append(&single_list.SingleNode{Data: 3})
	sl.Append(&single_list.SingleNode{Data: 4})
	sl.Append(&single_list.SingleNode{Data: 5})

	fmt.Println(sl.Length())
	sl.Display()
	fmt.Println(sl.Get(2))
	sl.Insert(3, &single_list.SingleNode{Data: 99})
	sl.Display()
	sl.Delete(5)
	sl.Display()
}
