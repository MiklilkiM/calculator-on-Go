package calculation

import "errors"

var (
	ErrInvalidExpression   = errors.New("invalid expression")    
	ErrDivisionByZero      = errors.New("division by zero")        
	ErrMismatchedParentheses = errors.New("mismatched parentheses")  
	ErrUnsupportedOperation = errors.New("unsupported operation")  
	ErrEmptyExpression     = errors.New("empty expression")       
	ErrOverflow            = errors.New("numeric overflow")      
)
