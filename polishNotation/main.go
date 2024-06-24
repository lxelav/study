package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Stack struct {
	operators []string
}

func (s *Stack) Push(operator string) {
	s.operators = append(s.operators, operator)
}

func (s *Stack) Pop() string {
	if len(s.operators) == 0 {
		return ""
	}

	operator := s.operators[len(s.operators)-1]
	s.operators = s.operators[:len(s.operators)-1]

	return operator
}

func (s *Stack) Peek() string {
	if len(s.operators) == 0 {
		return ""
	}

	operator := s.operators[len(s.operators)-1]
	return operator
}

func precedence(operator string) int {
	if operator == "+" || operator == "-" {
		return 1
	} else if operator == "*" || operator == "/" {
		return 2
	}
	return 0
}

func (s *Stack) IsEmpty() bool {
	return len(s.operators) == 0
}

func infixToPostfix(expression string) []string {
	resultPostfix := []string{}
	operators := Stack{}

	expressionLen := len(expression)
	temp := ""
	for i := 0; i < expressionLen; {
		char := rune(expression[i])
		if unicode.IsDigit(char) {
			for i < expressionLen && unicode.IsDigit(rune(expression[i])) {
				temp += string(expression[i])
				i++
			}
			resultPostfix = append(resultPostfix, temp)
			temp = ""
		} else if unicode.IsLetter(char) {
			for i < expressionLen && (unicode.IsDigit(rune(expression[i])) || unicode.IsLetter(rune(expression[i]))) {
				temp += string(expression[i])
				i++
			}
			resultPostfix = append(resultPostfix, temp)
			temp = ""
		} else if char == '+' || char == '-' || char == '/' || char == '*' {
			for !operators.IsEmpty() && precedence(operators.Peek()) >= precedence(string(char)) {
				resultPostfix = append(resultPostfix, operators.Pop())
			}
			operators.Push(string(char))
			i++
		} else if char == '(' {
			operators.Push(string(char))
			i++
		} else if char == ')' {
			for !operators.IsEmpty() && operators.Peek() != "(" {
				resultPostfix = append(resultPostfix, operators.Pop())
			}
			if !operators.IsEmpty() && operators.Peek() == "(" {
				operators.Pop()
			}
			i++
		} else {
			i++
		}
	}

	for !operators.IsEmpty() {
		resultPostfix = append(resultPostfix, operators.Pop())
	}

	return resultPostfix
}

func evaluatePostfix(expression []string) (int, error) {
	stack := &Stack{}

	for _, token := range expression {
		if unicode.IsDigit(rune(token[0])) || (len(token) > 1 && token[0] == '-' && unicode.IsDigit(rune(token[1]))) {
			// Если токен — число, конвертируем его и помещаем в стек
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", token)
			}
			stack.Push(num)
		} else {
			// Если токен — оператор, извлекаем два верхних числа из стека и выполняем операцию
			if stack.IsEmpty() {
				return 0, fmt.Errorf("insufficient values in the expression")
			}
			operand2 := stack.Pop()
			if stack.IsEmpty() {
				return 0, fmt.Errorf("insufficient values in the expression")
			}
			operand1 := stack.Pop()

			var result int
			switch token {
			case "+":
				result = operand1 + operand2
			case "-":
				result = operand1 - operand2
			case "*":
				result = operand1 * operand2
			case "/":
				result = operand1 / operand2
			default:
				return 0, fmt.Errorf("invalid operator: %s", token)
			}
			stack.Push(result)
		}
	}

	if stack.IsEmpty() {
		return 0, fmt.Errorf("no values in the stack")
	}
	result := stack.Pop()
	if !stack.IsEmpty() {
		return 0, fmt.Errorf("more values left in the stack")
	}
	return result, nil
}

func main() {
	expression := "2 + 3 * 4"
	postfix := infixToPostfix(expression)
	fmt.Println("Postfix expression:", postfix)
}
