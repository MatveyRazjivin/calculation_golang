package calculation

import (
	"strconv"
	"strings"
	"unicode"
)

func Calc(expression string) (float64, error) {
	part1, _ := stringToRPN(expression)
	return CalcRPN(part1)
}

func isOperator(token rune) bool {
	if token == '+' || token == '-' || token == '*' || token == '/' || token == '^' {
		return true
	}
	return false
}

func stringToRPN(expression string) (string, error) {
	var result string
	var stack []rune

	precedence := map[rune]int{
		'(': 0,
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'^': 3,
	}

	for _, token := range expression {
		if unicode.IsDigit(token) {
			result += string(token) + " "
		} else if token == '(' {
			stack = append(stack, '(')
		} else if token == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				result += string(stack[len(stack)-1]) + " "
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if isOperator(token) {
			if len(stack) > 0 && precedence[stack[len(stack)-1]] >= precedence[token] {
				result += string(stack[len(stack)-1]) + " "
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}

	for len(stack) > 0 {
		result += string(stack[len(stack)-1]) + " "
		stack = stack[:len(stack)-1]
	}

	return result, nil
}

func CalcRPN(expression string) (float64, error) {
	stack := []float64{}
	tokens := strings.Split(expression, " ")

	for _, token := range tokens {
		if token == "+" {
			if len(stack) < 2 {
				return 0, ErrIndexOfRange
			}
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			result := operand1 + operand2
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else if token == "-" {
			if len(stack) < 2 {
				return 0, ErrIndexOfRange
			}
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			result := operand1 - operand2
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else if token == "*" {
			if len(stack) < 2 {
				return 0, ErrIndexOfRange
			}
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			result := operand1 * operand2
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else if token == "/" {
			if len(stack) < 2 {
				return 0, ErrIndexOfRange
			}
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]

			if operand2 == 0 {
				return 0, ErrDivisionByZero
			}

			result := operand1 / operand2
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		} else {
			//number, _ := strconv.ParseFloat(token, 64)
			/*if err != nil {
				return 0, ErrInvalidToken
			}*/
			if number, err := strconv.ParseFloat(token, 64); err == nil {
				stack = append(stack, number)
			}
			//stack = append(stack, number)
		}
	}

	if len(stack) != 1 {
		return 0, ErrIndexOfRange
	}

	return stack[0], nil
}
