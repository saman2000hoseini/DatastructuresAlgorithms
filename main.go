package main

import (
	"DsAlgorithms/algorithm"
	"DsAlgorithms/ds"
)

func main()  {
	myList := ds.NewLinkedList()
	myList.PushBack(2)
	myList.PushBack(4)
	myList.AddFront(0)
	myList.AddNodeToIndex(3,5)
	//myList.Print()
	algorithm.Permute([]rune("hello"),0,4)
}
