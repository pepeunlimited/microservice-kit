package errorz

type Stack struct {
	stack []error
}

func (stack *Stack) Add(err error) {
	stack.stack = append(stack.stack, err)
}

func (stack *Stack) Pop() error {
	if len(stack.stack) > 0 {
		lastIndex := len(stack.stack) - 1
		err := stack.stack[lastIndex] //top element
		stack.stack = stack.stack[:lastIndex]
		return err
	}
	return nil
}

func NewErrorStack(errors []error) Stack {
	if errors == nil {
		errors = make([]error, 0)
	}
	return Stack{stack:errors}
}

func (stack Stack) IsEmpty() bool {
	return stack.stack == nil || len(stack.stack) == 0
}