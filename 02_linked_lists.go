package main

//------------------------------------------------------------------------------
// Problem 2.1

type listNode struct {
	val int
	next *listNode
}

type list struct {
	first *listNode
	last *listNode
}

func (node *listNode) getNext() *listNode {
	return node.next
}

func (lst *list) append(val int) *list {
	node := listNode{ val, nil }

	if (lst.first == nil) {
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
