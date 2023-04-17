package data_structures

import (
	"fmt"
	"testing"
)

/*
TestSimpledLinkedList

	input: 5, 3, 1, 2, 10, 12, 4, 8, 9
	expected :
		- head = 1
		- tail = 12
		- length of list node = 9
*/
func TestSimpleLinkedList(t *testing.T) {
	//Create linked list
	list := LinkedList{}

	//add value
	input := []int{5, 3, 1, 2, 10, 12, 4, 8, 9}

	for idx, in := range input {
		list.AddNode(fmt.Sprintf("%d", idx), in)
	}

	//value head should be 1
	expectedHeadValue := 1
	if list.Head.val == expectedHeadValue {
		t.Logf("head value is %d", list.Head.val)
	} else {
		t.Errorf("got %d, expected %d", list.Head.val, expectedHeadValue)
	}

	//value tail should be 12
	expectedTailValue := 12
	if list.Tail.val == expectedTailValue {
		t.Logf("tail value is %d", list.Tail.val)
	} else {
		t.Errorf("tail value got %d, expected %d", list.Tail.val, expectedTailValue)
	}

	//total node should be 9
	expectedTotalNode := 9
	if len(list.ListNode) == expectedTotalNode {
		t.Logf("total node is %d", len(list.ListNode))
	} else {
		t.Errorf("total node got %d, expected %d", len(list.ListNode), expectedTotalNode)
	}
}
