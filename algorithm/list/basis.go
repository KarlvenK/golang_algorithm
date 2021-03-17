package list

//insert e  after at
func (l *List) insert(e, at *Node) *Node {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.length++
	return e
}

//insertValue is a convenient wrapper for insert &Node{Value : v}
func (l *List) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}
