// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply takes two numbers and returns the result of the product of both parameters
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of the division of both parameters
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Sqrt takes a number and returns its square root. If the given number is negative an error is returned
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("square root of negative value is not allowed")
	}
	return math.Sqrt(a), nil
}
