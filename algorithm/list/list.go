package list

type listInterface interface {

}

type listNode struct {
	Value interface{}
	next, prev *node
}

type List struct {
	root   listNode
	length int
}

func New() * {

}
