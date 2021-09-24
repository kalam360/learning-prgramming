package main

import (
	"fmt"
	"container/list"
)


func list_check() {
	var intList list.List
	intList.PushBack(11)
	intList.PushBack(23)
	intList.PushBack(45)

	for element := intList.Front(); element != nil; element = element.Next() {

		fmt.Println(element.Value.(int))

	}
}