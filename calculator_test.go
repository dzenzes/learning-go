package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

type testCase struct {
	a, b        float64
	want        float64
	name        string
	errExpected bool
}

func TestAdd(t *testing.T) {
	t.Parallel()

	testCases := []testCase{
		{a: 2, b: 2, want: 4, name: "Normal case with two positive numbers"},
		{a: 2, b: -1, want: 1, name: "Positive and negative number"},
		{a: 5, b: 0, want: 5, name: "Adding zero to a number"},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.0000001) {
			t.Errorf("%s: Add(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	for i := 0; i < 100; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if !closeEnough(want, got, 0.0000001) {
			t.Errorf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}

}

func TestSubtract(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, b: 2, want: 2, name: "Normal example with two positive numbers"},
		{a: 1, b: 3, want: -2, name: "Subtracting a bigger number from the first number to get a negative result"},
		{a: -1, b: -1, want: 0, name: "Subtracting two negative numbers"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.0000001) {
			t.Errorf("%s: Subtract(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, b: 2, want: 8, name: "Normal example with positive numbers"},
		{a: 1, b: -1, want: -1, name: "Multiply with one negative number"},
		{a: -2, b: -2, want: 4, name: "Two negative numbers result in a positive result"},
		{a: 4, b: 0, want: 0, name: "Using a zero results in a zero"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if !closeEnough(tc.want, got, 0.0000001) {
			t.Errorf("%s: Multiply(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []testCase{
		{a: 4, b: 2, want: 2, name: "Normal example with positive numbers"},
		{a: 1, b: -1, want: -1, name: "Divide with one negative number"},
		{a: -2, b: -4, want: 0.5, name: "Two negative numbers result in positive result"},
		{a: 4, b: 0, errExpected: true, name: "Shouldn'T be possible to divide by zero"},
	}

	for _, tc := range testCases {

		got, err := calculator.Divide(tc.a, tc.b)

		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status: %v", tc.a, tc.b, errReceived)
		}

		if !tc.errExpected && !closeEnough(tc.want, got, 0.0000001) {
			t.Errorf("%s: Divide(%f, %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
		// The error value was as you expected, and the data value also matched expectation (pass)

	}
}
