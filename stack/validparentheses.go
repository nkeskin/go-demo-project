package stack

func isValid(input string) bool {
	var s Stack
	for _, v := range input {
		if !s.IsEmpty() {
			temp := s.Peek().(rune)
			if checkCompatibility(temp, v) {
				s.Pop()
			} else {
				s.Push(v)
			}
		} else {
			s.Push(v)
		}
	}
	return s.IsEmpty()
}

func checkCompatibility(r1 rune, r2 rune) bool {
	if r1 == '(' && r2 == ')' {
		return true
	}
	if r1 == '{' && r2 == '}' {
		return true
	}
	if r1 == '[' && r2 == ']' {
		return true
	}
	return false
}
