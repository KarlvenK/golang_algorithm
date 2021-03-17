package queue

type Queue struct {
	lenght int
	data   []interface{}
}

func New() Queue {
	return new(Queue).Init()
}

func (q *Queue) Init() Queue {
	q.data = make([]interface{}, 0)
	q.lenght = 0
	return *q
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Empty() bool {
	return q.lenght == 0
}

func (q *Queue) Push(v interface{}) {
	q.data = append(q.data, v)
	q.lenght++
}

func (q *Queue) Pop() interface{} {
	defer func() {
		q.data = q.data[1:q.lenght]
		q.lenght--
	}()
	return q.data[0]
}
