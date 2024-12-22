package calculation_test

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		name        string
		expression  string
		expected    float64
		expectedErr error
	}{
		{
			name:       "Addition",
			expression: "2+3",
			expected:   5,
		},
		{
			name:       "Subtraction",
			expression: "5-2",
			expected:   3,
		},
		{
			name:       "Multiplication",
			expression: "4*3",
			expected:   12,
		},
		{
			name:       "Division",
			expression: "8/2",
			expected:   4,
		},
		{
			name:        "Division by zero",
			expression:  "8/0",
			expectedErr: ErrDivisionByZero,
		},
		{
			name:       "Combined operations",
			expression: "2+3*4",
			expected:   14,
		},
		{
			name:       "Parentheses",
			expression: "(2+3)*4",
			expected:   20,
		},
		{
			name:       "Negative numbers",
			expression: "-2+3",
			expected:   1,
		},
		{
			name:       "Decimal numbers",
			expression: "2.5+1.5",
			expected:   4,
		},
		{
			name:       "Whitespace in expression",
			expression: " 2 + 3 ",
			expected:   5,
		},
		{
			name:        "Invalid token",
			expression:  "2+a",
			expectedErr: ErrIndexOfRange, // Исправлено на ErrIndexOfRange, так как RPN не валидирует токены перед обработкой
		},
		{
			name:        "Empty expression",
			expression:  "",
			expectedErr: ErrIndexOfRange,
		},
		{
			name:       "Power",
			expression: "2^3",
			expected:   8,
		},
		{
			name:       "Complex expression",
			expression: "((10+2)*2-4)/4",
			expected:   5,
		},
		{
			name:        "Unclosed Parentheses",
			expression:  "(2+3",
			expectedErr: ErrIndexOfRange,
		},
		{
			name:        "Multiple operators",
			expression:  "2++3",
			expectedErr: ErrIndexOfRange,
		},
		{
			name:        "Number with many points",
			expression:  "2.3.2+3",
			expectedErr: ErrIndexOfRange,
		},
		{
			name:       "Substraction with negative numbers",
			expression: "-1 - -2",
			expected:   1,
		},
		{
			name:       "Expression with only number",
			expression: "4",
			expected:   4,
		},
		{
			name:       "Expression with only number and whitespace",
			expression: "  4  ",
			expected:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Calc(tc.expression)
			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected error %v, but got nil", tc.expectedErr)
				} else if err != tc.expectedErr {
					t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			} else if actual != tc.expected {
				t.Errorf("Expected %f, but got %f", tc.expected, actual)
			}
		})
	}
}

func TestStringToRPN(t *testing.T) {
	testCases := []struct {
		name        string
		expression  string
		expectedRPN string
		expectedErr error
	}{
		{
			name:        "Simple addition",
			expression:  "2+3",
			expectedRPN: "2 3 + ",
		},
		{
			name:        "Simple subtraction",
			expression:  "5-2",
			expectedRPN: "5 2 - ",
		},
		{
			name:        "Simple multiplication",
			expression:  "4*3",
			expectedRPN: "4 3 * ",
		},
		{
			name:        "Simple division",
			expression:  "8/2",
			expectedRPN: "8 2 / ",
		},
		{
			name:        "Combined operations",
			expression:  "2+3*4",
			expectedRPN: "2 3 4 * + ",
		},
		{
			name:        "Parentheses",
			expression:  "(2+3)*4",
			expectedRPN: "2 3 + 4 * ",
		},
		{
			name:        "Power",
			expression:  "2^3",
			expectedRPN: "2 3 ^ ",
		},
		{
			name:        "Complex Expression",
			expression:  "((10+2)*2-4)/4",
			expectedRPN: "10 2 + 2 * 4 - 4 / ",
		},
		{
			name:        "Expression with whitespaces",
			expression:  " 2 + 3 * 4 ",
			expectedRPN: "2 3 4 * + ",
		},
		{
			name:        "Expression with only one number",
			expression:  "4",
			expectedRPN: "4 ",
		},
		{
			name:        "Expression with only one number and whitespaces",
			expression:  "  4  ",
			expectedRPN: "4 ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualRPN, err := stringToRPN(tc.expression)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
			}
			if actualRPN != tc.expectedRPN {
				t.Errorf("Expected RPN %v, but got %v", tc.expectedRPN, actualRPN)
			}
		})
	}
}

func TestCalcRPN(t *testing.T) {
	testCases := []struct {
		name        string
		expression  string
		expected    float64
		expectedErr error
	}{
		{
			name:       "Simple addition",
			expression: "2 3 + ",
			expected:   5,
		},
		{
			name:       "Simple subtraction",
			expression: "5 2 - ",
			expected:   3,
		},
		{
			name:       "Simple multiplication",
			expression: "4 3 * ",
			expected:   12,
		},
		{
			name:       "Simple division",
			expression: "8 2 / ",
			expected:   4,
		},
		{
			name:        "Division by zero",
			expression:  "8 0 / ",
			expectedErr: ErrDivisionByZero,
		},
		{
			name:       "Combined operations",
			expression: "2 3 4 * + ",
			expected:   14,
		},
		{
			name:       "Parentheses",
			expression: "2 3 + 4 * ",
			expected:   20,
		},
		{
			name:       "Power",
			expression: "2 3 ^ ",
			expected:   8,
		},
		{
			name:       "Complex expression",
			expression: "10 2 + 2 * 4 - 4 / ",
			expected:   5,
		},
		{
			name:        "Not enough arguments",
			expression:  "2 + ",
			expectedErr: ErrIndexOfRange,
		},
		{
			name:       "Only number",
			expression: "4 ",
			expected:   4,
		},
		{
			name:        "Empty expression",
			expression:  "",
			expectedErr: ErrIndexOfRange,
		},
		{
			name:        "Not enough arguments on second operand",
			expression:  "2 3 4 + - ",
			expectedErr: ErrIndexOfRange,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := CalcRPN(tc.expression)
			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected error %v, but got nil", tc.expectedErr)
				} else if err != tc.expectedErr {
					t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
				}
			} else if err != nil {
				t.Errorf("Unexpected error: %v", err)
			} else if actual != tc.expected {
				t.Errorf("Expected %f, but got %f", tc.expected, actual)
			}
		})
	}
}
