package ds

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	len int
	head *node
	tail *node
}

func newNode(data int) *node {
	return &node{data,nil}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{0,nil,nil}
}

func (myList *LinkedList) PushBack(data int) {
	if myList.head==nil {
		myList.head = newNode(data)
		myList.tail = myList.head
		return
	}
	myList.tail.next = newNode(data)
	myList.tail = myList.tail.next
	myList.len++
}

func (myList *LinkedList) AddFront(data int) {
	if myList.head==nil {
		myList.head = newNode(data)
		myList.tail = myList.head
		return
	}
	first := newNode(data)
	first.next = myList.head
	myList.head = first
	myList.len++
}

func (myList *LinkedList) Print() {
	curr := myList.head
	for curr!=nil{
		fmt.Print(curr.data," ")
		curr = curr.next
	}
	fmt.Println()
}

func (myList *LinkedList) RemoveData(data int) {
	if myList.head==nil {
		return
	}
	if myList.head.data==data {
		if myList.head == myList.tail {
			myList.tail = nil
		}
		myList.head = myList.head.next
	}
	curr := myList.head
	var last *node
	for curr != nil{
		if curr.data==data {
			if curr==myList.tail {
				myList.tail = last
			}
			last.next = curr.next
		}
		last = curr
		curr = curr.next
	}
	myList.len--
}

func (myList *LinkedList) RemoveIndex(index int) {
	if myList.head==nil{
		return
	}
	if index == 0{
		myList.head = myList.head.next
		if index==myList.len {
			myList.tail = nil
		}
		myList.len--
		return
	}
	if index <= myList.len {
		curr := myList.head
		for i := 0; i < index-1; i++ {
			curr = curr.next
		}
		if index==myList.len {
			myList.tail = curr
		}
		curr.next = curr.next.next
		myList.len--
	}
}

func (myList *LinkedList) AddNodeToIndex(index int,data int){
	if index==myList.len+1{
		myList.PushBack(data)
	}else if index == 0{
		myList.AddFront(data)
	}else if index<=myList.len {
		cur := myList.head
		for i :=1;i<index;i++{
			cur = cur.next
		}
		newN := newNode(data)
		newN.next = cur.next
		cur.next = newN
	}
}

func (myList *LinkedList) Get(index int) int{
	if index>myList.len {
		return -1
	}
	cur:= myList.head
	for i:=0; i<index; i++ {
		cur=cur.next
	}
	return cur.data
}