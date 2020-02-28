package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

// func main() {
// 	//输入，暂未构造测试用例
// 	list1 := &Node{value: 1}
// 	list2 := &Node{value: 2}

// 	node1 := list1
// 	node2 := list2
// 	newList := &Node{}
// 	newNode := newList

// 	for node1 != nil && node2 != nil {
// 		if node1.value > node2.value {
// 			newNode.value = node2.value
// 			node2 = node2.next
// 		} else {
// 			newNode.value = node1.value
// 			node1 = node1.next
// 		}
// 		newNode.next = &Node{}
// 		newNode = newNode.next
// 	}
// 	for node1 != nil {
// 		newNode.value = node1.value
// 		node1 = node1.next
// 		newNode.next = &Node{}
// 		newNode = newNode.next
// 	}
// 	for node2 != nil {
// 		newNode.value = node2.value
// 		node2 = node2.next
// 		newNode.next = &Node{}
// 		newNode = newNode.next
// 	}

// 	//已经构造升序链表，反转链表即可，并且由于之前未处理最后一个多余节点的关系需要删除反转后的第一个节点
// 	newHead := solve(newList).next
// 	//test
// 	node := newHead
// 	for node != nil {
// 		fmt.Println(node.value)
// 		node = node.next
// 	}
// }

// func solve(node *Node) *Node {
// 	if node.next == nil {
// 		return node
// 	}
// 	newHead := solve(node.next)
// 	node.next.next = node
// 	node.next = nil
// 	return newHead
// }

func main() {
	head := &Node{value: 1}
	result1 := is(head)

	tmp := head
	k := 3
	for i := 1; i < k-1 && tmp != nil; i++ {
		tmp = tmp.next
	}

	result2 := false
	if tmp != nil && tmp.next != nil {
		tmp.next = tmp.next.next
		result2 = is(head)
	}

	fmt.Println(result1, result2)
}
func is(head *Node) bool {
	flag := find(head)
	head2 := solve(flag)
	for head != nil && head2 != nil {
		if head.value != head2.value {
			return false
		}
		head = head.next
		head2 = head2.next
	}
	return true
}
func find(node *Node) *Node {
	n1, n2 := node, node
	for n2 != nil && n2.next != nil {
		n1 = n1.next
		n2 = n2.next.next
	}
	return n1

}
func solve(node *Node) *Node {
	if node.next == nil {
		return node
	}
	newHead := solve(node.next)
	node.next.next = node
	node.next = nil
	return newHead
}
