package data_structures

type Node struct {
	val int

	left  *Node
	right *Node
}

type LinkedList struct {
	Head     *Node
	Tail     *Node
	ListNode map[string]*Node
}

func (l *LinkedList) Add(newNode *Node) {
	if l.Head == nil && l.Tail == nil {
		l.Head = newNode
		l.Tail = newNode
	} else if newNode.val < l.Head.val {
		newNode.right = l.Head
		l.Head.left = newNode
		l.Head = newNode
	} else if newNode.val > l.Tail.val {
		newNode.left = l.Tail
		l.Tail.right = newNode
		l.Tail = newNode
	} else {
		if newNode.val-l.Head.val < l.Tail.val-newNode.val {
			//start from head
			temp := l.Head.right

			for temp != nil {
				if newNode.val > temp.val {
					//next
					temp = temp.right
				} else {
					newNode.right = temp
					newNode.left = temp.left
					temp.left.right = newNode
					temp.left = newNode
					break
				}
			}
		} else {
			//start from tail
			temp := l.Tail.left

			for temp != nil {
				if newNode.val < temp.val {
					temp = temp.left
				} else {
					newNode.left = temp
					newNode.right = temp.right
					temp.right.left = newNode
					temp.right = newNode
					break
				}
			}
		}
	}
}

func (l *LinkedList) AddNode(key string, val int) {
	if l.ListNode == nil {
		l.ListNode = make(map[string]*Node)
	}

	newNode := Node{val: val}
	l.ListNode[key] = &newNode

	l.Add(&newNode)
}

func (l *LinkedList) UpdateNode(key string, val int) {
	//create new node
	prevNode := l.ListNode[key]

	//delete node
	l.Delete(prevNode)

	//add new node
	newNode := Node{val: val}
	l.Add(&newNode)
	l.ListNode[key] = &newNode
}

func (l *LinkedList) Delete(n *Node) {
	if n == nil {
		return
	}

	if l.Head == l.Tail && l.Head == n {
		l.Tail = nil
		l.Head = nil
	} else if n == l.Tail {
		l.Tail = n.left
		l.Tail.right = nil
	} else if n == l.Head {
		l.Head = n.right
		l.Head.left = nil
	} else {
		n.right.left = n.left
		n.left.right = n.right
	}
}
