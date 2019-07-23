package main

import (
	"errors"
	"fmt"
	"reflect"
)

type ElementType interface{}

type LNode struct {
	elem ElementType
	next *LNode
}

type LinkedList struct {
	Head *LNode
}

func NewLNode(elem ElementType, next *LNode) *LNode {
	return &LNode{elem, next}
}

func NewLinkedList() *LinkedList {
	head := NewLNode(0, nil)
	return &LinkedList{head}
}

func (list *LinkedList) IsEmpty() bool {
	return list.Head.next == nil
}

func (list *LinkedList) Length() int {
	return int(reflect.ValueOf(list.Head.elem).Int())
}

func (list *LinkedList) Append(node *LNode) {
	current := list.Head
	for {
		if current.next == nil {
			break
		}
		current = current.next
	}
	current.next = node
	list.sizeInc()
}

func (list *LinkedList) Prepend(node *LNode) {
	current := list.Head
	node.next = current.next
	current.next = node
	list.sizeInc()
}

func (list *LinkedList) Find(x ElementType) (*LNode, bool) {
	empty := list.IsEmpty()
	if empty {
		fmt.Println("This is an empty list")
		return nil, false
	}
	current := list.Head
	for current.next != nil {
		if current.elem == x {
			return current, true
		}
		current = current.next
	}
	return nil, false
}

func (list *LinkedList) Remove(x ElementType) error {
	empty := list.IsEmpty()
	if empty {
		return errors.New("This is an empty list")
	}
	current := list.Head
	for current.next != nil {
		if current.next.elem == x {
			current.next = current.next.next
			list.sizeDec()
			return nil
		}
		current = current.next
	}
	return nil
}

func (list *LinkedList) sizeInc() {
	v := int(reflect.ValueOf((*list.Head).elem).Int())
	list.Head.elem = v + 1
}

func (list *LinkedList) sizeDec() {
	v := int(reflect.ValueOf((*list.Head).elem).Int())
	list.Head.elem = v - 1
}

func (list *LinkedList) PrintList() {
	empty := list.IsEmpty()
	if empty {
		fmt.Println("This is an empty list")
		return
	}
	current := list.Head.next
	fmt.Println("The elements is:")
	i := 0
	for ; ; i++ {
		if current.next == nil {
			break
		}
		fmt.Printf("INode%d ,value:%v -> ", i, current.elem)
		current = current.next
	}
	fmt.Printf("Node%d value:%v", i+1, current.elem)
	return
}

func main() {
	linklist := NewLinkedList()
	linklist.Append(&LNode{3, nil})
	linklist.Append(&LNode{4, nil})
	linklist.Append(&LNode{5, nil})
	linklist.Prepend(&LNode{7, nil})
	linklist.PrintList()
}
