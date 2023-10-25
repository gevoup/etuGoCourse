package linkedList

import (
	"errors"
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

func (l *LinkedList) Pop() error {
	newNode := l.head
	if newNode == nil {
		return errors.New("list is empty")
	}
	if newNode.next == nil {
		l.head = nil
		return nil
	}
	for newNode.next.next != nil {
		newNode = newNode.next
	}
	newNode.next = nil
	return nil
}

func (l *LinkedList) At(pos int) (int, error) {
	newNode := l.head
	if newNode == nil {
		return 0, errors.New("list is empty")
	}
	for i := 0; i < pos; i++ {
		newNode = newNode.next
	}
	return newNode.val, nil
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

func (l *LinkedList) DeleteFrom(pos int) error {
	if l.head == nil {
		return errors.New("list is empty")
	}
	if pos == 0 {
		if l.head.next == nil {
			l.head = nil
		} else {
			l.head = l.head.next
		}
		return nil
	}
	newNode := l.head
	for i := 1; i < pos; i++ {
		newNode = newNode.next
	}
	newNode.next = newNode.next.next
	return nil
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