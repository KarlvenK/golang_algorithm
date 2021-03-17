package list

type Node struct {
	Value interface{} // value of this Node

	next, prev *Node // next and previous pointers in the doubly linked list of Nodes

	list *List //the list which this Node belongs
}

type List struct {
	root   Node
	length int
}

// Init initializes or clean list l
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.length = 0
	return l
}

func New() *List {
	return new(List).Init()
}

func (l *List) Len() int {
	return l.length
}

func (l *List) Front() *Node {
	if l.length == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Back() *Node {
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

func (n *Node) Prev() *Node {
	if n.list != nil && n.prev != &n.list.root {
		return n.prev
	}
	return nil
}

func (n *Node) Next() *Node {
	if n.list != nil && n.next != &n.list.root {
		return n.next
	}
	return nil
}

func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

//InsertBefore insert a new node e with value v before mark and return e
//if mark is not a node of l, then return immediately
func (l *List) InsertBefore(v interface{}, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

func (l *List) InsertAfter(v interface{}, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// remove removes e from its list, returns e
func (l *List) remove(e *Node) *Node {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.length--
	return e
}

func (l *List) Remove(e *Node) interface{} {
	if e.list == l {
		l.remove(e)
	}
	return e.Value
}

// move moves e to next to at and return e
func (l *List) move(e, at *Node) *Node {
	if e == at {
		return e
	}
	e.prev.next = e.next
	e.next.prev = e.prev

	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e

	return e
}

func (l *List) MoveToFront(e *Node) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

func (l *List) MoveToBack(e *Node) {
	if e.list != l || l.root.prev == e {
		return
	}
	l.move(e, l.root.prev)
}

func (l *List) MoveBefore(e, mark *Node) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark.prev)
}

func (l *List) MoveAfter(e, mark *Node) {
	if e.list != l || e == mark || mark.list != l {
		return
	}
	l.move(e, mark)
}

func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.next {
		l.insertValue(e.Value, l.root.prev)
	}
}

func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.prev {
		l.insertValue(e.Value, l.root.prev)
	}
}
