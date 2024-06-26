package polishNotation

import (
	. "project_dict/stack"
	"unicode"
)

func priority(operation string) int {
	if operation == "+" || operation == "-" {
		return 1
	} else if operation == "*" || operation == "/" {
		return 2
	}
	return 0
}

func InfixToPrefix(expression string) []string {
	result := []string{}
	stackOperations := Stack{}

	lenExpression := len(expression)
	for i := 0; i < lenExpression; {
		char := rune(expression[i])
		if unicode.IsDigit(char) {
			start := i
			for i < lenExpression && (unicode.IsDigit(rune(expression[i]))) {
				i++
			}
			result = append(result, expression[start:i])
		} else if unicode.IsLetter(char) {
			start := i
			for i < lenExpression && (unicode.IsDigit(rune(expression[i])) || unicode.IsLetter(rune(expression[i]))) {
				i++
			}
			result = append(result, expression[start:i])
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			for !stackOperations.IsEmpty() && (priority(stackOperations.Peek().(string)) >= priority(string(char))) {
				result = append(result, stackOperations.Pop().(string))
			}
			stackOperations.Push(string(char))
			i++
		} else if char == '(' {
			stackOperations.Push(string(char))
			i++
		} else if char == ')' {
			for !stackOperations.IsEmpty() && stackOperations.Peek() != "(" {
				result = append(result, stackOperations.Pop().(string))
			}
			stackOperations.Pop()
			i++
		} else {
			i++
		}
	}

	for !stackOperations.IsEmpty() {
		result = append(result, stackOperations.Pop().(string))
	}

	return result
}
