package linkedList

import (
	"fmt"
	"strconv"
)

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head *node
}

func New(q int) *LinkedList {
	newNode := &node{}
	newLinkedList := &LinkedList{head: newNode}
	for i := 1; i < q; i++ {
		next := &node{}
		newNode.next = next
		newNode = next
	}
	return newLinkedList
}

func (l *LinkedList) Add(val int) {
	newNode := l.head
	for newNode.next != nil {
		newNode = newNode.next
	}
	newNode.next = &node{val: val}
}

func (l *LinkedList) Pop() {
	newNode := l.head
	if newNode.next == nil {
		l.head = nil
		return
	}
	for newNode.next.next != nil {
		newNode = newNode.next
	}
	newNode.next = nil
}

func (l *LinkedList) At(pos int) int {
	newNode := l.head
	for i := 0; i < pos; i++ {
		newNode = newNode.next
	}
	return newNode.val
}

func (l *LinkedList) Size() int {
	newNode := l.head
	if newNode == nil {
		return 0
	}
	count := 1
	for newNode.next != nil {
		count += 1
		newNode = newNode.next
	}
	return count
}

func (l *LinkedList) DeleteFrom(pos int) {
	if pos == 0 {
		if l.head.next == nil {
			l.head = nil
		} else {
			l.head = l.head.next
		}
		return
	}
	newNode := l.head
	for i := 1; i < pos; i++ {
		newNode = newNode.next
	}
	newNode.next = newNode.next.next
}

func (l *LinkedList) UpdateAt(pos, val int) {
	if pos == 0 {
		l.head.val = val
		return
	}
	newNode := l.head
	for i := 1; i <= pos; i++ {
		newNode = newNode.next
	}
	newNode.val = val
}

func NewFromSlice(s []int) *LinkedList {
	newNode := &node{val: s[0]}
	l := &LinkedList{head: newNode}
	for i := 1; i < len(s); i++ {
		newNode.next = &node{val: s[i]}
		newNode = newNode.next
	}
	return l
}

func (l *LinkedList) print() {
	res := ""
	newNode := l.head
	for newNode.next != nil {
		res += strconv.Itoa(newNode.val) + " -> "
		newNode = newNode.next
	}
	res += strconv.Itoa(newNode.val)
	fmt.Println(res)
}