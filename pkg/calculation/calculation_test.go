package calculation

import (
	"errors"
	"testing"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   float64
		expectErr  error
	}{
		{"Addition", "3 + 5", 8, nil},
		{"Subtraction", "10 - 2", 8, nil},
		{"Multiplication", "4 * 2", 8, nil},
		{"Division", "16 / 2", 8, nil},

		{"Complex Expression", "3 + 5 * 2", 13, nil},
		{"Parentheses Multiplication", "(1 + 2) * 3", 9, nil},
		{"Parentheses Division", "(1 + 2) / 3", 1, nil},
		{"Nested Parentheses Multiplication", "2 * (3 + 5)", 16, nil},
		{"Nested Parentheses Subtraction", "2 * (3 + 5) - 4", 12, nil},
		{"Nested Parentheses Division", "2 * (3 + 5) / 4", 4, nil},

		{"Division by Zero", "2 / 0", 0, ErrDivisionByZero},
		{"Mismatched Parentheses", "2 * (3 + 5", 0, ErrMismatchedParentheses},
		{"Unsupported Character", "2 + a", 0, ErrInvalidExpression},
		{"Empty Expression", "", 0, ErrEmptyExpression},

		// Ошибочные примеры
		{"Invalid Expression Missing Operand", "3 +", 0, ErrInvalidExpression},
		{"Invalid Expression Missing Operator", "* 3 + 5", 0, ErrInvalidExpression},
		{"Multiple Operators in a Row", "3 ++ 2", 0, ErrInvalidExpression},
		{"Overflow Error", "1e309", 0, ErrOverflow},
		{"Negative Division by Zero", "-4 / 0", 0, ErrDivisionByZero},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Calc(tt.expression)

			if result != tt.expected {
				t.Errorf("Result mismatch: got %v, want %v", result, tt.expected)
			}

			if !errors.Is(err, tt.expectErr) {
				t.Errorf("Error mismatch: got %v, want %v", err, tt.expectErr)
			}
		})
	}
}
