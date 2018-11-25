package stack

// EvalPostfix 함수는 연산자 우선순위가 동일한 후위 표기식 계산 결과를 반환
func EvalPostfix(exp string) int {
	operation := map[string]func(op1 int, op2 int) int{
		"+": func(op1 int, op2 int) int {
			return op1 + op2
		},
		"-": func(op1 int, op2 int) int {
			return op1 - op2
		},
		"*": func(op1 int, op2 int) int {
			return op1 * op2
		},
		"/": func(op1 int, op2 int) int {
			return op1 / op2
		},
	}
	stack := New(10)

	for _, char := range exp {
		if operation[string(char)] == nil {
			stack.Push(int(char - '0'))
			continue
		}
		oper2 := stack.Pop().(int)
		oper1 := stack.Pop().(int)
		result := operation[string(char)](oper1, oper2)
		stack.Push(result)
	}
	return stack.Pop().(int)
}
