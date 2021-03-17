package list

type Interface interface {
}

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

func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

func (l *List) PushBack(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

//InsertBefore insert a new node e with value v before mark and return e
//of makr is not a node of l, then return immediately
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
