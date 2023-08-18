package stack

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str interface{}) {
	*s = append(*s, str)
}

func (s *Stack) Pop() interface{} {
	elem := (*s)[len(*s)-1]
	*s = (*s)[:(len(*s) - 1)]
	return elem
}

func (s *Stack) Peek() interface{} {
	return (*s)[len(*s)-1]
}
