package calculator

import (
	"errors"
)

func Calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("operator tidak dikenal")
	}
}
