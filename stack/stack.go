package stack

type Stack struct {
	Values []interface{}
}

func NewStack(args ...interface{}) (*Stack, error) {
	result := Stack{}
	for _, el := range args {
		result.Values = append(result.Values, el)
	}
	return &Stack{}, nil
}

func (stack *Stack) Push(value interface{}) {
	stack.Values = append(stack.Values, value)
}

func (stack *Stack) Peek() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.Values[len(stack.Values)-1]
}

func (stack *Stack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	value := stack.Values[len(stack.Values)-1]
	stack.Values = stack.Values[:len(stack.Values)-1]

	return value
}

func (stack *Stack) IsEmpty() bool {
	return len(stack.Values) == 0
}
