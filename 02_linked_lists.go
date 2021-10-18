package main

import (
	"fmt"
	"strconv"
	"strings"
)

//------------------------------------------------------------------------------
// Initial

type listNode struct {
	val  int
	next *listNode
}

type list struct {
	first *listNode
	last  *listNode
}

func (node *listNode) getNext() *listNode {
	return node.next
}

func (lst *list) append(val int) *list {
	node := listNode{val, nil}

	if lst.first == nil {
		lst.first = &node
	} else {
		lst.last.next = &node
	}
	lst.last = &node

	return lst
}

func toList(values []int) *list {
	lst := new(list)
	for _, n := range values {
		lst.append(n)
	}
	return lst
}

func printList(lst *list) {
	var sb strings.Builder
	for node := lst.first; node != nil; node = node.next {
		sb.WriteString(strconv.Itoa(node.val))
		sb.WriteString(", ")
	}
	fmt.Println(sb.String())
}

//------------------------------------------------------------------------------
// Problem 2.1

func main() {
	arr := []int{1, 2, 3, 4}
	list := toList(arr)
	printList(list)
}
