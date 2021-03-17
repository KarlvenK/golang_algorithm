package stack

type Stack struct {
	length int
	data   []interface{}
}

func (s *Stack) Init() Stack {
	s.length = 0
	s.data = make([]interface{}, 0)
	return *s
}

func New() Stack {
	return new(Stack).Init()
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v)
	s.length++
}

func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	defer func() {
		s.data = s.data[:s.length-1]
		s.length--
	}()
	return s.data[s.length-1]
}

func (s *Stack) Top() interface{} {
	return s.data[s.length-1]
}

func (s *Stack) Empty() bool {
	return s.length == 0
}
